CREATE SCHEMA IF NOT EXISTS phishing_quest;

CREATE TABLE IF NOT EXISTS phishing_quest.users (
                                                    id UUID PRIMARY KEY,
                                                    username VARCHAR(255) NOT NULL UNIQUE,
                                                    email VARCHAR(255) NOT NULL UNIQUE,
                                                    password_hash VARCHAR(255) NOT NULL,
                                                    total_score INT DEFAULT 0,
                                                    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                                    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS phishing_quest.categories (
                                                         id UUID PRIMARY KEY,
                                                         category_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS phishing_quest.questions (
                                                        id UUID PRIMARY KEY,
                                                        category_id UUID NOT NULL REFERENCES phishing_quest.categories(id),
                                                        question_text TEXT NOT NULL,
                                                        correct_answer VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS phishing_quest.answers (
                                                      id UUID PRIMARY KEY,
                                                      question_id UUID NOT NULL REFERENCES phishing_quest.questions(id),
                                                      answer_text VARCHAR(255) NOT NULL,
                                                      is_correct BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS phishing_quest.user_answers (
                                                           user_answer_id UUID PRIMARY KEY,
                                                           user_id UUID NOT NULL REFERENCES phishing_quest.users(id),
                                                           question_id UUID NOT NULL REFERENCES phishing_quest.questions(id),
                                                           answer_id UUID NOT NULL REFERENCES phishing_quest.answers(id),
                                                           is_correct BOOLEAN,
                                                           timestamp TIMESTAMPTZ DEFAULT NOW(),
                                                           ai_rating INT CHECK (ai_rating BETWEEN 1 AND 5) DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS phishing_quest.user_scores (
                                                          id UUID PRIMARY KEY,
                                                          user_id UUID NOT NULL REFERENCES phishing_quest.users(id),
                                                          score INT NOT NULL,
                                                          timestamp TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_answers_user_id ON phishing_quest.user_answers(user_id);
CREATE INDEX IF NOT EXISTS idx_user_answers_question_id ON phishing_quest.user_answers(question_id);
CREATE INDEX IF NOT EXISTS idx_user_answers_answer_id ON phishing_quest.user_answers(answer_id);

COMMENT ON COLUMN phishing_quest.user_answers.ai_rating IS 'Avaliação da IA de 1 a 5: 1-Péssimo, 2-Ruim, 3-Mediano, 4-Bom, 5-Excelente';
