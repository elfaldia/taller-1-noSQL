package App

func main() {
	conexion, err := connection.connectToDataBase()
	if err != nil{
		println("error")
	}
}
