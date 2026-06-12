package tasks

type Wrapper struct {
    Currency Currency `json:"currency"`
}

type Currency struct {
    RUR string `json:"RUR"`
    USD string `json:"USD"`
    EUR string `json:"EUR"`
    JPY string `json:"JPY"`
}