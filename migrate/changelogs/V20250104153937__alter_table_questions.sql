ALTER TABLE phishing_quest.questions
DROP COLUMN IF EXISTS correct_answer;

COMMENT ON TABLE phishing_quest.questions IS 'Tabela de perguntas associadas às categorias';
COMMENT ON TABLE phishing_quest.answers IS 'Tabela de respostas associadas às perguntas';
