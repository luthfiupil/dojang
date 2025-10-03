-- =======================================
-- Seed Data for Lookup Tables
-- =======================================

-- Roles
INSERT INTO roles (role_name) VALUES
                                  ('athlete'),
                                  ('coach'),
                                  ('admin');

-- Genders
INSERT INTO genders (gender_name) VALUES
                                      ('male'),
                                      ('female'),
                                      ('other');

-- Belt Ranks
INSERT INTO belt_ranks (rank_name, rank_order) VALUES
                                                   ('White', 1),
                                                   ('Yellow', 2),
                                                   ('Green', 3),
                                                   ('Blue', 4),
                                                   ('Red', 5),
                                                   ('Black 1st Dan', 6),
                                                   ('Black 2nd Dan', 7),
                                                   ('Black 3rd Dan', 8);

-- Weight Classes (example Olympic style)
INSERT INTO weight_classes (class_name, min_weight, max_weight) VALUES
                                                                    ('Flyweight', 0, 58.00),
                                                                    ('Featherweight', 58.01, 68.00),
                                                                    ('Welterweight', 68.01, 80.00),
                                                                    ('Heavyweight', 80.01, 999.00);

-- Coach Ranks
INSERT INTO coach_ranks (rank_name) VALUES
                                        ('1st Dan'),
                                        ('2nd Dan'),
                                        ('3rd Dan'),
                                        ('4th Dan'),
                                        ('5th Dan');

-- Attendance Status
INSERT INTO attendance_status (status_name) VALUES
                                                ('present'),
                                                ('absent'),
                                                ('late');

-- Test Results
INSERT INTO test_results (result_name) VALUES
                                           ('passed'),
                                           ('failed');

-- Tournament Results
INSERT INTO tournament_results (result_name) VALUES
                                                 ('Gold'),
                                                 ('Silver'),
                                                 ('Bronze'),
                                                 ('Eliminated');
