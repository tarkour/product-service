package db

type Product struct {
	id           int
	uniq_id      string
	brand_name   string //full name, no need to make it short
	category_id  int    //glasses, pullover, etc. - one to one соединение
	main_picture string //link to storage with pics - parse all pics from tg-channel, then manually compare
	color        string //should it be choosable colours or any? quantity int8 #most of products are uniq, but sometimes there are more then just one
	created_at   string //datetime
}

// массив для одной картинки, которая через product_id служит для many to one
type Product_image struct {
	id         int
	product_id int
	url        string
}
type Category struct {
	id   int
	name string
}
