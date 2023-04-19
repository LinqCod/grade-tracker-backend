CREATE TYPE "user_role" AS ENUM ('student', 'teacher', 'admin');

CREATE TABLE "users" (
                            "id" bigserial,
                            "email" varchar NOT NULL,
                            "password" varchar NOT NULL,
                            "first_name" varchar NOT NULL,
                            "second_name" varchar NOT NULL,
                            "patronymic" varchar NOT NULL,
                            "role" user_role,
                            PRIMARY KEY ("id", "email")
);

CREATE TABLE "groups" (
                            "id" bigserial PRIMARY KEY,
                            "title" varchar NOT NULL
);

CREATE TABLE "students" (
                            group_id bigint NOT NULL REFERENCES groups(id),
                            PRIMARY KEY ("id", "email")
) INHERITS(users);

CREATE TABLE "teachers" (
                            PRIMARY KEY ("id", "email")
) INHERITS(users);

CREATE TABLE "admins" (
                            PRIMARY KEY ("id", "email")
) INHERITS(users);

CREATE TABLE "subjects" (
                          "id" bigserial PRIMARY KEY,
                          "title" varchar NOT NULL
);

CREATE TABLE "groups_subjects" (
                                   "group_id" bigint NOT NULL REFERENCES groups(id),
                                   "subject_id" bigint NOT NULL REFERENCES subjects(id),
                                   CONSTRAINT groups_subjects_pkey PRIMARY KEY (group_id, subject_id)
);

CREATE TABLE "teachers_subjects" (
                                   "teacher_id" bigint NOT NULL REFERENCES teachers(id),
                                   "subject_id" bigint NOT NULL REFERENCES subjects(id),
                                   CONSTRAINT teachers_subjects_pkey PRIMARY KEY (teacher_id, subject_id)
);

CREATE TABLE "teachers_groups" (
                                     "teacher_id" bigint NOT NULL REFERENCES teachers(id),
                                     "group_id" bigint NOT NULL REFERENCES groups(id),
                                     CONSTRAINT teachers_groups_pkey PRIMARY KEY (teacher_id, group_id)
);

CREATE TYPE "day_of_week" AS ENUM ('ПОНЕДЕЛЬНИК', 'ВТОРНИК', 'СРЕДА', 'ЧЕТВЕРГ', 'ПЯТНИЦА', 'СУББОТА');

CREATE TABLE "schedules" (
                            "id" bigserial PRIMARY KEY,
                            "start_time" varchar NOT NULL,
                            "end_time" varchar NOT NULL,
                            "day_of_week" day_of_week
);

CREATE TABLE "schedules_subjects" (
                                      "schedule_id" bigint NOT NULL REFERENCES schedules(id),
                                      "subject_id" bigint NOT NULL REFERENCES subjects(id),
                                      CONSTRAINT schedules_subjects_pkey PRIMARY KEY (schedule_id, subject_id)
);

CREATE TABLE "schedules_groups" (
                                      "schedule_id" bigint NOT NULL REFERENCES schedules(id),
                                      "group_id" bigint NOT NULL REFERENCES groups(id),
                                      CONSTRAINT schedules_groups_pkey PRIMARY KEY (schedule_id, group_id)
);

CREATE TABLE "schedules_teachers" (
                                    "schedule_id" bigint NOT NULL REFERENCES schedules(id),
                                    "teacher_id" bigint NOT NULL REFERENCES teachers(id),
                                    CONSTRAINT schedules_teachers_pkey PRIMARY KEY (schedule_id, teacher_id)
);