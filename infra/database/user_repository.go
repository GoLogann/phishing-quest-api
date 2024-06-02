package database

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *domain.User) error {
    _, err := r.db.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", user.Username, user.Email, user.PasswordHash)
    return err
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
    row := r.db.QueryRow("SELECT user_id, username, email, password_hash, xp, total_score FROM users WHERE user_id = $1", id)
    user := &domain.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.XP, &user.TotalScore)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
    row := r.db.QueryRow("SELECT user_id, username, email, password_hash, xp, total_score FROM users WHERE email = $1", email)
    user := &domain.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.XP, &user.TotalScore)
    if err != nil {
        return nil, err
    }
    return user, nil
}