-- Criação do schema phishing_quest, se ainda não existir
CREATE SCHEMA IF NOT EXISTS phishing_quest;

-- Criação da tabela users
CREATE TABLE IF NOT EXISTS phishing_quest.users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    xp INT DEFAULT 0,
    total_score INT DEFAULT 0
    );

-- Criação da tabela categories
CREATE TABLE IF NOT EXISTS phishing_quest.categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL UNIQUE
);

-- Criação da tabela questions
CREATE TABLE IF NOT EXISTS phishing_quest.questions (
    question_id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES phishing_quest.categories(category_id),
    question_text TEXT NOT NULL,
    correct_answer VARCHAR(255) NOT NULL
    );

-- Criação da tabela answers
CREATE TABLE IF NOT EXISTS phishing_quest.answers (
    answer_id SERIAL PRIMARY KEY,
    question_id INT NOT NULL REFERENCES phishing_quest.questions(question_id),
    answer_text VARCHAR(255) NOT NULL,
    is_correct BOOLEAN NOT NULL
    );

-- Criação da tabela user_answers
CREATE TABLE IF NOT EXISTS phishing_quest.user_answers (
    user_answer_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES phishing_quest.users(user_id),
    question_id INT NOT NULL REFERENCES phishing_quest.questions(question_id),
    answer_id INT NOT NULL REFERENCES phishing_quest.answers(answer_id),
    is_correct BOOLEAN,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ai_rating INT CHECK (ai_rating BETWEEN 1 AND 5) DEFAULT NULL
    );

-- Criação da tabela user_scores
CREATE TABLE IF NOT EXISTS phishing_quest.user_scores (
    user_score_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES phishing_quest.users(user_id),
    score INT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Adicionando índices para otimização
CREATE INDEX IF NOT EXISTS idx_user_answers_user_id ON phishing_quest.user_answers(user_id);
CREATE INDEX IF NOT EXISTS idx_user_answers_question_id ON phishing_quest.user_answers(question_id);
CREATE INDEX IF NOT EXISTS idx_user_answers_answer_id ON phishing_quest.user_answers(answer_id);

-- Adicionando comentário para clarificar o propósito da coluna ai_rating
COMMENT ON COLUMN phishing_quest.user_answers.ai_rating IS 'Avaliação da IA de 1 a 5: 1-Péssimo, 2-Ruim, 3-Mediano, 4-Bom, 5-Excelente';
