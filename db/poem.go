package db

type Poem struct {
	Id int
	Title string
	Author string
	Dynasty string
	Content string
}

func (p *Poem) Insert() bool  {
	stmtInsert, err := db.Prepare("INSERT INTO poem(title,author,dynasty,content) VALUES (?,?,?,?)")
	if checkError(err) {
		return  false
	}
	_,err = stmtInsert.Exec(&p.Title, &p.Author, &p.Dynasty, &p.Content)

	if checkError(err) {
		return  false
	}

	return true
}

func (p *Poem) Save()  {
	p.Insert()
}