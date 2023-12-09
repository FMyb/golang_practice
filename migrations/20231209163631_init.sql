-- +goose Up
-- +goose StatementBegin
create table if not exists hospitals
(
    hospital_id uuid primary key,
    name        text unique
);

create table if not exists patients
(
    patient_id uuid primary key,
    name       text,
    age        int
);

create table if not exists registered_patients
(
    patient_id uuid not null references patients(patient_id),
    hospital_id uuid not null references hospitals(hospital_id),
    register_time timestamp,
    discard_time timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists hospitals;
drop table if exists patients;
-- +goose StatementEnd
