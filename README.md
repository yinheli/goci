goci
====
> 这是我目前找到的最好的一个go语言开发oralce数据库的oci接口。
>
> 总体实现的比较好，还是标准的。
>
> 推荐使用。

oci for go. It's goci!

例子：
        package main

        import (
                "database/sql"
                _ "goci"  // 根据实际部署情况修改
                "os"
                "log"
        )

        func main() {
                // 为log添加短文件名,方便查看行数
                log.SetFlags(log.Lshortfile | log.LstdFlags)

                log.Println("Oracle Driver example")

                os.Setenv("NLS_LANG", "")
                dsn := os.Getenv("ORACLE_DSN") // 把用户名/口令@SID  定义到此环境变量中
                if dsn == "" {
        		t.Fatal("To run tests, set the ORACLE_DSN environment variable.")
                }
        	db, _ := driver.Open(dsn)

                rows, err := db.Query("select 3.14, 'foo' from dual")
                if err != nil {
                        log.Fatal(err)
                }
                defer db.Close()

                for rows.Next() {
                        var f1 float64
                        var f2 string
                        rows.Scan(&f1, &f2)
                        log.Println(f1, f2) // 3.14 foo
                }
                rows.Close()

                // 先删表,再建表
                db.Exec("drop table sdata")
                db.Exec("create table sdata(name varchar2(256))")

                db.Exec("insert into sdata values('中文')")
                db.Exec("insert into sdata values('1234567890ABCabc!@#$%^&*()_+')")

                rows, err = db.Query("select * from sdata")
                if err != nil {
                        log.Fatal(err)
                }

                for rows.Next() {
                        var name string
                        rows.Scan(&name)
                        log.Printf("Name = %s, len=%d", name, len(name))
                }
                rows.Close()
        }


