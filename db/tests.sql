create table tests (
    id         uuid primary key not null,
    solver     varchar(255)     not null,
    boardstr   char(81)         not null,
    numclues   int              not null,
    difficulty varchar(255)     not null,
    solvetime  real             not null
)