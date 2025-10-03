-- =======================================
-- Taekwondo Dojang Management (Initial Schema)
-- =======================================

-- Lookup Tables
CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,
                       role_name VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE genders (
                         id SERIAL PRIMARY KEY,
                         gender_name VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE belt_ranks (
                            id SERIAL PRIMARY KEY,
                            rank_name VARCHAR(50) UNIQUE NOT NULL,
                            rank_order INT UNIQUE NOT NULL
);

CREATE TABLE weight_classes (
                                id SERIAL PRIMARY KEY,
                                class_name VARCHAR(50) UNIQUE NOT NULL,
                                min_weight DECIMAL(5,2),
                                max_weight DECIMAL(5,2)
);

CREATE TABLE coach_ranks (
                             id SERIAL PRIMARY KEY,
                             rank_name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE attendance_status (
                                   id SERIAL PRIMARY KEY,
                                   status_name VARCHAR(20) UNIQUE NOT NULL -- present, absent, late
);

CREATE TABLE test_results (
                              id SERIAL PRIMARY KEY,
                              result_name VARCHAR(20) UNIQUE NOT NULL -- passed, failed
);

CREATE TABLE tournament_results (
                                    id SERIAL PRIMARY KEY,
                                    result_name VARCHAR(50) UNIQUE NOT NULL -- Gold, Silver, Bronze, Eliminated
);

-- Core Tables
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       full_name VARCHAR(100) NOT NULL,
                       date_of_birth DATE,
                       gender_id INT REFERENCES genders(id),
                       phone VARCHAR(20),
                       email VARCHAR(100) UNIQUE,
                       address TEXT,
                       role_id INT REFERENCES roles(id) NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE athletes (
                          user_id INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                          belt_rank_id INT REFERENCES belt_ranks(id),
                          weight_class_id INT REFERENCES weight_classes(id),
                          join_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE coaches (
                         user_id INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                         specialization VARCHAR(100),
                         rank_id INT REFERENCES coach_ranks(id),
                         hire_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE attendance (
                            id SERIAL PRIMARY KEY,
                            user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                            status_id INT REFERENCES attendance_status(id),
                            check_in TIMESTAMP,
                            check_out TIMESTAMP,
                            remarks TEXT,
                            created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE training_sessions (
                                   id SERIAL PRIMARY KEY,
                                   coach_id INT REFERENCES coaches(user_id) ON DELETE SET NULL,
                                   session_date DATE NOT NULL,
                                   topic VARCHAR(200),
                                   notes TEXT
);

CREATE TABLE training_participation (
                                        id SERIAL PRIMARY KEY,
                                        session_id INT REFERENCES training_sessions(id) ON DELETE CASCADE,
                                        athlete_id INT REFERENCES athletes(user_id) ON DELETE CASCADE,
                                        attended BOOLEAN DEFAULT FALSE,
                                        remarks TEXT
);

CREATE TABLE belt_tests (
                            id SERIAL PRIMARY KEY,
                            athlete_id INT REFERENCES athletes(user_id) ON DELETE CASCADE,
                            test_date DATE NOT NULL,
                            current_rank_id INT REFERENCES belt_ranks(id),
                            next_rank_id INT REFERENCES belt_ranks(id),
                            result_id INT REFERENCES test_results(id),
                            examiner_id INT REFERENCES coaches(user_id) ON DELETE SET NULL
);

CREATE TABLE tournaments (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(200) NOT NULL,
                             location VARCHAR(200),
                             start_date DATE NOT NULL,
                             end_date DATE,
                             organizer VARCHAR(100)
);

CREATE TABLE tournament_participation (
                                          id SERIAL PRIMARY KEY,
                                          tournament_id INT REFERENCES tournaments(id) ON DELETE CASCADE,
                                          athlete_id INT REFERENCES athletes(user_id) ON DELETE CASCADE,
                                          weight_class_id INT REFERENCES weight_classes(id),
                                          result_id INT REFERENCES tournament_results(id),
                                          notes TEXT
);
