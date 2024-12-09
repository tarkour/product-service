struct Product{ 
    id int
    uniq_id string
    brand_name string //full name, no need to make it short
    category_id int //glasses, pullover, etc. - one to one соединение
    main_picture string //link to storage with pics - parse all pics from tg-channel, then manually compare
    color string //should it be choosable colours or any? quantity int8 #most of products are uniq, but sometimes there are more then just one
    created_at datetime 
 }


 // массив для одной картинки, которая через product_id служит для many to one
 struct Product_image{
    id int
    product_id int
    url string
 }


 struct Category{
    id int
    name string
 }