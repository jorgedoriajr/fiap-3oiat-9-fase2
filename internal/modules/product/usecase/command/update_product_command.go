package command

type UpdateProductCommand struct {
	Number      int
	Name        *string
	Description *string
	Category    *string
	Menu        *bool
	ImgPath     *string
	Ingredients []Ingredient
}

func NewUpdateProductCommand(
	Number int,
	Name *string,
	Description *string,
	Category *string,
	Menu *bool,
	Ingredients []Ingredient,
	ImgPath *string,
) *UpdateProductCommand {

	cmd := &UpdateProductCommand{
		Number:      Number,
		Name:        Name,
		Description: Description,
		Category:    Category,
		Menu:        Menu,
		ImgPath:     ImgPath,
		Ingredients: Ingredients,
	}
	return cmd
}
