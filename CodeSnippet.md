### go myql select解析

```golang
func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
    if err != nil{
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM user WHERE gid = 1")
    if err != nil{
        log.Fatalln(err)
    }
    defer rows.Close()
    
    
    cols, err := rows.Columns()
    if err != nil{
        log.Fatalln(err)
    }
    fmt.Println(cols)
    vals := make([][]byte, len(cols))
    scans := make([]interface{}, len(cols))
    
    for i := range vals{
        scans[i] = &vals[i]
    }
    
    var results []map[string]string
    
    for rows.Next(){
        err = rows.Scan(scans...)
        if err != nil{
            log.Fatalln(err)
        }
    
        row := make(map[string]string)
        for k, v := range vals{
            key := cols[k]
            row[key] = string(v)
        }
        results = append(results, row)
    }
    
    for k, v :=range results{
        fmt.Println(k, v)
    }
}
```