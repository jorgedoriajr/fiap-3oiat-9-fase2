package command

type UpdateProductCommand struct {
	Number      int
	Name        *string
	Description *string
	Category    *string
	Menu        *bool
	ImgPath     *string
}

func NewUpdateProductCommand(number int, name *string, description *string, category *string, menu *bool, imgPath *string,
) UpdateProductCommand {
	return UpdateProductCommand{
		Number: number, Name: name, Description: description, Category: category, Menu: menu, ImgPath: imgPath,
	}
}
