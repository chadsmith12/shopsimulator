package customers

type Customer struct {
    id int 
}

func New(id int) Customer {
    return Customer{ id: id }
}
