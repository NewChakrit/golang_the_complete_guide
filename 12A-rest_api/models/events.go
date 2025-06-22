package models

import (
	"example.com/rest_api/db"
)

var events []Event

func (e *Event) Save() error {
	query := `INSERT INTO events(name ,description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)` // todo "?" เพื่อแทนค่า column ต่างๆ ใน table ตามลำดับ

	stmt, err := db.DB.Prepare(query) // todo prepare ข้อมูลนั้นจะถูกเก็บไว้ในหน่วยความจำ และสามารถนำไปใช้ซ้ำได้
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // todo Exec(args ...any) รับ args หลายตัวเพื่อรองรับค่า query
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // auto increment ID ให้
	if err != nil {
		return err
	}

	e.ID = id

	events = append(events, Event{})

	return nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // ปิดเมื่อ func ทำงานเสร็จ

	var events []Event

	for rows.Next() { // จะ loop เอาค่า row ถัดไปเรื่อยๆจนหมด
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (e Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?  WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events  WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

/* todo Note
DB.Exec() vs. DB.Query()
ก่อนอื่น เรามาทำความเข้าใจความแตกต่างหลักระหว่าง Exec() และ Query() ก่อน:

DB.Exec(): ใช้เมื่อคุณต้องการ จัดการฐานข้อมูลหรือข้อมูลในฐานข้อมูล แต่ ไม่ได้คาดหวังผลลัพธ์กลับมา ครับ ตัวอย่างเช่น:

สร้างตาราง (CREATE TABLE)
แทรกข้อมูล (INSERT INTO)
อัปเดตข้อมูล (UPDATE)
ลบข้อมูล (DELETE FROM)
คำสั่งเหล่านี้จะส่งคืน sql.Result ซึ่งบอกจำนวนแถวที่ได้รับผลกระทบหรือ ID ที่แทรกเข้าไป (ถ้ามี) แต่ไม่ส่งคืนชุดข้อมูล
DB.Query(): ใช้เมื่อคุณต้องการ ดึงข้อมูลจากฐานข้อมูล และ คาดหวังผลลัพธ์กลับมา ครับ ตัวอย่างเช่น:

เลือกข้อมูล (SELECT)
คำสั่งนี้จะส่งคืน *sql.Rows ซึ่งคุณสามารถวนลูปเพื่ออ่านข้อมูลแต่ละแถวได้
คุณสามารถส่งคำสั่ง SQL ทั้งหมดโดยตรงผ่าน Exec() หรือ Query() ได้เลยครับ ไม่จำเป็นต้องใช้ Prepare() เสมอไป

ข้อดีของ Prepare() (และข้อควรระวัง)
Prepare() ทำหน้าที่ "เตรียม" คำสั่ง SQL

Prepare() + stmt.Exec() หรือ stmt.Query():
ขั้นตอนแรกคือการเรียกใช้ DB.Prepare(query string) ซึ่งจะคืนค่า *sql.Stmt (Statement object)
จากนั้น คุณสามารถเรียกใช้เมธอด Exec() หรือ Query() บน *sql.Stmt นั้นได้ (stmt.Exec() หรือ stmt.Query())
แล้วข้อดีของการใช้ Prepare() คืออะไร?

Prepare() จะทำการวิเคราะห์ (parse) และปรับปรุงประสิทธิภาพ (optimize) คำสั่ง SQL ของคุณล่วงหน้า หนึ่งครั้ง เมื่อคุณเรียกใช้ Prepare()

ประสิทธิภาพที่ดีขึ้น (Potential Performance Benefit):
ถ้าคำสั่ง SQL เดียวกันถูกเรียกใช้ หลายครั้ง (โดยอาจมีข้อมูลที่แตกต่างกันในส่วนของ placeholder เช่น ? หรือ $1 ในคำสั่ง INSERT) การใช้ Prepare() สามารถนำไปสู่ประสิทธิภาพที่ดีขึ้นได้
เหตุผลก็คือ ฐานข้อมูลไม่ต้องเสียเวลาวิเคราะห์และปรับปรุงคำสั่งเดิมซ้ำๆ ทุกครั้งที่รัน เพราะมันถูก "เตรียม" ไว้แล้ว
ข้อควรระวัง:

ข้อดีด้านประสิทธิภาพนี้จะเกิดขึ้นได้ก็ต่อเมื่อ *sql.Stmt ที่ถูกเตรียมไว้ ไม่ได้ถูกปิด (stmt.Close()) ระหว่างการเรียกใช้หลายครั้ง
หากคุณเรียก stmt.Close() ทันทีหลังจาก stmt.Exec() หรือ stmt.Query() ในแต่ละครั้ง คุณก็จะไม่ได้ประโยชน์อะไรเลยจากการใช้ Prepare() เพราะฐานข้อมูลจะต้องเตรียมคำสั่งนั้นใหม่ทุกครั้งอยู่ดี
ในบริบทของแอปพลิเคชันนี้
อย่างที่ระบุไว้ในเนื้อหาที่คุณให้มา:

ในแอปพลิเคชันนี้ คุณกำลังเรียกใช้ stmt.Close() ทันทีหลังจากเรียกใช้ stmt.Exec()
ดังนั้น ในสถานการณ์นี้ การใช้ Prepare() จึงไม่มีข้อได้เปรียบด้านประสิทธิภาพเลย ครับ คุณสามารถใช้ DB.Exec() โดยตรงได้เช่นเดียวกับที่ใช้สร้างตาราง
สรุป
ผู้สอนได้ตัดสินใจรวมวิธีการ Prepare() เข้ามาในคอร์สนี้ก็เพื่อ แสดงให้เห็นถึงวิธีที่แตกต่างกันในการใช้แพ็กเกจ database/sql ครับ
ถึงแม้ว่าในกรณีเฉพาะของแอปพลิเคชันนี้จะไม่ได้ให้ประโยชน์ด้านประสิทธิภาพที่ชัดเจน แต่การรู้จักและเข้าใจ Prepare()
เป็นสิ่งสำคัญสำหรับการเขียนโค้ด Go ที่โต้ตอบกับฐานข้อมูลอย่างมีประสิทธิภาพในสถานการณ์จริงที่คำสั่งเดียวกันถูกรันซ้ำๆ ครับ


*/
