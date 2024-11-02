package appError

type Duplicated struct {
    Message *string
}

var _ error = (*Duplicated)(nil)

func (e Duplicated) Error() string {
    if e.Message == nil {
        return "Duplicated"
    }
    return *e.Message
}
