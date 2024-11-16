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

type Process struct {
    Message *string
}

var _ error = (*Process)(nil)

func (e Process) Error() string {
    if e.Message == nil {
        return "Processing"
    }
    return *e.Message
}
