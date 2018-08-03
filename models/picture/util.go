package models

import "fmt"

type util interface {
	 GetName()
}

func (p * Picture ) GetName()  {
	fmt.Println(p.Name)
}

func (p * Picture ) getIsPublic()  {
	fmt.Println(p.IsPublic)
}
