package main

import (
    //_ "github.com/eaciit/dbox/dbc/mongo"
    //"github.com/eaciit/dbox"
    //"github.com/eaciit/orm"
    "github.com/eaciit/toolkit"
    "github.com/eaciit/config"
    "github.com/eaciit/knot/knot.v1"
    "path/filepath"
    "os"
)

func App() *knot.App{
    app := knot.NewApp("procura")
    wd, _ := os.Getwd()
    app.ViewsPath = filepath.Join(wd,"views")
    app.LayoutTemplate="_layout.html"
    app.Static("static",filepath.Join(wd,"assets"))
    d := new(Dashboard)
    app.Register(d)
    //app.Register(new(Dashboard))
    app.DefaultOutputType = knot.OutputHtml
    return app
}

var log *toolkit.LogEngine

func main() {
    log, _ := toolkit.NewLog(true,false,"","", "")

    configpath, _ := os.Getwd()
    configpath = filepath.Join(configpath,"..","config","app.json")
    econfig := config.SetConfigFile(configpath)
    if econfig!=nil {
        log.Error("Error loading config file " + econfig.Error())
    }

    port := int(config.GetDefault("port",9100).(float64))
    app := App()
    knot.StartApp(app, toolkit.Sprintf("localhost:%d", port))
}

/*
func conn() dbox.IConnection{
    conn, _ := dbox.NewConnection("mongo",&dbox.ConnectionInfo{"localhost:27123","ectest","","",nil})
    conn.Connect()
    return conn
}

func DB() *orm.DataContext{
    return orm.New(conn())
}
*/