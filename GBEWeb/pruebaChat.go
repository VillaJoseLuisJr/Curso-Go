package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pruebas_go?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Crear tabla si no existe
	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME
		);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1 - Insertar nuevo usuario")
		fmt.Println("2 - Listar todos los usuarios")
		fmt.Println("3 - Eliminar usuario por ID")
		fmt.Println("4 - Salir")
		fmt.Print("Elegí una opción: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var username, password string
			fmt.Print("Nombre de usuario: ")
			fmt.Scanln(&username)
			fmt.Print("Contraseña: ")
			fmt.Scanln(&password)

			_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, time.Now())
			if err != nil {
				log.Println("Error al insertar:", err)
			} else {
				fmt.Println("✅ Usuario insertado correctamente.")
			}

		case 2:
			rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
			if err != nil {
				log.Println("Error al consultar usuarios:", err)
				continue
			}
			defer rows.Close()

			fmt.Println("\n🧾 Lista de usuarios:")
			for rows.Next() {
				var u User
				if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt); err != nil {
					log.Println("Error al leer fila:", err)
					continue
				}
				fmt.Printf("- ID: %d | Usuario: %s | Contraseña: %s | Creado: %s\n", u.ID, u.Username, u.Password, u.CreatedAt.Format("2006-01-02 15:04:05"))
			}

		case 3:
			var id int
			fmt.Print("ID del usuario a eliminar: ")
			fmt.Scanln(&id)

			result, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
			if err != nil {
				log.Println("Error al eliminar:", err)
				continue
			}
			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 0 {
				fmt.Println(" No se encontró un usuario con ese ID.")
			} else {
				fmt.Println(" Usuario eliminado correctamente.")
			}

		case 4:
			fmt.Println(" Saliendo...")
			return

		default:
			fmt.Println(" Opción inválida. Probá de nuevo.")
		}
	}
}
