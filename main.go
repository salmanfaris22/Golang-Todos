package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type StoreBlog struct {
	stor []Blog
}

type Blog struct {
	head   string
	body   string
	privet bool
}

func (s *StoreBlog) addBlog(head, body string) {

	newBlog := Blog{head: head, body: body, privet: true}
	s.stor = append(s.stor, newBlog)
	fmt.Println("Bloge Added 🦚")
	// fmt.Println(s.stor)
}

func (s *StoreBlog) listBlogs() {
	if len(s.stor) == 0 {
		fmt.Println("No Bloge Availible 💔")
	} else {
		for i, v := range s.stor {
			privet := "Public"
			if v.privet {
				privet = "Privite"
			}
			fmt.Printf("id:%d: 👉🏻title is := %s , \n  🤜🏻%s \n  🥶this Bloge Is %s \n", i+1, v.head, v.body, privet)
		}
	}

}

func (s *StoreBlog) setStaus(i int, b string) error {
	if i > len(s.stor) || i <= 0 {
		return errors.New("eror")
	}
	if b == "private" {
		s.stor[i-1].privet = true
		fmt.Println("Staus Changed 🦚")
		return nil
	} else {
		s.stor[i-1].privet = false
		fmt.Println("Staus Changed 🦚")
		return nil

	}

}

func (s *StoreBlog) editsBlog(i int) error {
	scanner := bufio.NewScanner(os.Stdin)
	if i > len(s.stor) || i <= 0 {
		return errors.New("eror")
	}

	fmt.Println(" 🦋What You Want Change head Or body Or both")
	scanner.Scan()
	get := scanner.Text()

	switch get {
	case "head":
		fmt.Println("🪴 Enter Your New Blogs")
		scanner.Scan()
		title := scanner.Text()
		s.stor[i-1].head = title
		fmt.Println("Bloge Edited 🦚")
	case "body":
		fmt.Println("🪴 Enter Your New Blogs")
		scanner.Scan()
		body := scanner.Text()
		s.stor[i-1].body = body
		fmt.Println("Bloge Edited 🦚")
	case "both":
		fmt.Println("🪴 Enter Your New Blogs")
		scanner.Scan()
		title := scanner.Text()
		fmt.Println("🪴 Enter Your New Blogs")
		scanner.Scan()
		body := scanner.Text()
		s.stor[i-1].head = title
		s.stor[i-1].body = body
		fmt.Println("Bloge Edited 🦚")
	default:
		fmt.Println("Enter valide string")

	}

	return nil
}

func (s *StoreBlog) deliteBloge(i int) error {
	if i > len(s.stor) || i <= 0 {
		return errors.New("Erre")
	} else {
		s.stor = append(s.stor[:i-1], s.stor[i:]...)
		fmt.Println("Bloge Delited")
	}

	return nil
}

func main() {
	fmt.Println("Welcome To Bloge Apps")
	store := StoreBlog{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("_____________________")
		fmt.Println("Bloge Option:")
		fmt.Println("1 : ADD Blog 🫡")
		fmt.Println("2 : List Blog 🫠")
		fmt.Println("3 : Set Privet Or Public 🤯")
		fmt.Println("4 : Edit Blog 🙌🏻")
		fmt.Println("5 : Delite Blog 👌🏻")
		fmt.Println("6: Exit")
		fmt.Println("_____________________")
		fmt.Println("🫵🏻 Enter Your Option : ")
		scanner.Scan()
		input := scanner.Text()
		newinput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalide Number ")
			continue
		}
		switch newinput {
		case 1:
			fmt.Println(" 🦋Enter Your Blog Title")
			scanner.Scan()
			title := scanner.Text()
			fmt.Println("🪴 Enter Your Blogs")
			scanner.Scan()
			body := scanner.Text()
			store.addBlog(title, body)

		case 2:
			fmt.Println("Listing Bloge🦊 ")
			store.listBlogs()
		case 3:
			fmt.Println("Enter The ID Of Bloge🦊 ")
			scanner.Scan()
			input := scanner.Text()
			index, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Pleas Enter valid NUmebr")
			}

			fmt.Println("public or private")
			scanner.Scan()
			status := scanner.Text()

			if status == "public" || status == "private" {
				err := store.setStaus(index, status)
				if err != nil {
					fmt.Println("Enter Your Error")
				}
			} else {
				fmt.Println("Enter Valide ")

			}
		case 4:
			fmt.Println("Which On is Your Want Edits enter ID :")
			scanner.Scan()
			input, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Ivalid Id")
			}
			er := store.editsBlog(input)
			if er != nil {
				fmt.Println("Enter Your Error Id :")
			}
		case 5:
			fmt.Println("Which On is Your Want Delite enter ID :")
			scanner.Scan()
			input, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Ivalid Id")
			}
			er := store.deliteBloge(input)
			if er != nil {
				fmt.Println("Ivalid Id")
			}
		case 6:
			fmt.Println("Exited")
			return
		default:

			continue

		}

	}

}
