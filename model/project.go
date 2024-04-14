import model
import "time"

type Project struct {
    Id          int
    Title       string
    Description string
    OwnerID     int // ID do usuário que criou o projeto
    Members     []int // IDs dos usuários que são membros do projeto
    Tasks       []Task // Tarefas associadas ao projeto
    CreatedAt   time.Time
}