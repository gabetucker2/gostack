![Banner](../images/gostack_Smaller.png)

 <h1>Installation</h1>

 1. Edit your **go.mod** file
   
    *([or add one if you haven't already](https://go.dev/doc/tutorial/create-module))*
 2. Add `github.com/gabetucker2/gostack [release version]` to **go.mod**'s **require** clause

 <img src="../images/releases.png" width="20%" style="margin-left:10%; margin-bottom:15px"/>
 <img src="../images/requirements.png" width="40%" style="margin-left:10%"/>

 3. Type `go mod tidy` in your terminal
 
    *(ensuring that **go.mod** is a direct child of your current directory)*

 4. Add `. "github.com/gabetucker2/gostack"` to the **imports** of every file in which you would like to use ***gostack***

 <img src="../images/imports.png" width="30%" style="margin-left:10%; margin-bottom:15px"/>

---
 [> Return to glossary](../README.md)