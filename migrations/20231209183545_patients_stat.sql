-- +goose Up
-- +goose StatementBegin
create or replace function patients_stat(_from timestamp, _to timestamp)
    returns table
            (
                id                 uuid,
                name               text,
                number_of_patients int,
                average_age        float,
                number_of_register int,
                number_of_discard  int
            )
as
$$
begin
    if _from > _to then
        return;
    end if;

    drop table if exists ages_stat;
    drop table if exists patients_stat;
    drop table if exists register_stat;
    drop table if exists discard_stat;

    create temp table ages_stat
    (
        hospital_id uuid,
        average_age float
    );

    create temp table patients_stat
    (
        hospital_id        uuid,
        number_of_patients int
    );
    create temp table register_stat
    (
        hospital_id        uuid,
        number_of_register int
    );
    create temp table discard_stat
    (
        hospital_id       uuid,
        number_of_discard int
    );

    insert into ages_stat (select r.hospital_id, avg(p.age)
                           from registered_patients r
                                    natural join patients p
                           where _from <= r.register_time and r.register_time <= _to
                              or _from <= r.discard_time and r.discard_time <= _to
                           group by r.hospital_id);

    insert into patients_stat (select r.hospital_id, count(*)
                               from registered_patients r
                               where _from <= r.register_time and r.register_time <= _to
                                  or _from <= r.discard_time and r.discard_time <= _to
                               group by r.hospital_id);

    insert into register_stat (select r.hospital_id, count(*)
                               from registered_patients r
                               where _from <= r.register_time
                                 and r.register_time <= _to
                               group by r.hospital_id);

    insert into discard_stat (select r.hospital_id, count(*)
                              from registered_patients r
                              where _from <= r.discard_time
                                and r.discard_time <= _to
                              group by r.hospital_id);
    return query
        select hospitals.hospital_id,
               hospitals.name,
               coalesce(p.number_of_patients, 0),
               coalesce(a.average_age, 0),
               coalesce(r.number_of_register, 0),
               coalesce(d.number_of_discard, 0)
        from hospitals
                 left join ages_stat a using (hospital_id)
                 left join patients_stat p using (hospital_id)
                 left join register_stat r using (hospital_id)
                 left join discard_stat d using (hospital_id);
end;
$$ language plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function if exists patients_stat;
-- +goose StatementEnd
