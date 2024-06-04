-- Создание таблицы клиентов
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(100),
    address TEXT
);

-- Создание таблицы отделов
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- Создание таблицы ролей пользователей
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);

-- Создание таблицы сотрудников
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(100),
    department_id INT REFERENCES Departments(id),
    password_hash TEXT NOT NULL,
    role_id INT REFERENCES Roles(id),
    position VARCHAR(100)
);

-- Создание таблицы кабинетов
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    room_number VARCHAR(10) NOT NULL,
    department_id INT REFERENCES Departments(id)
);

-- Связующая таблица между сотрудниками и кабинетами (многие ко многим)
CREATE TABLE employee_rooms (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES Employees(id),
    room_id INT REFERENCES Rooms(id)
);

-- Создание таблицы расходных материалов в кабинетах
CREATE TABLE room_storage (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES Rooms(id),
    item_name VARCHAR(100) NOT NULL,
    quantity INT NOT NULL
);

-- Создание таблицы заявок на посещение врача
CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES Clients(id),
    employee_id INT REFERENCES Employees(id),
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(50)
);

-- Создание таблицы посещений врача
CREATE TABLE medical_sessions (
    id SERIAL PRIMARY KEY,
    session_date DATE NOT NULL,
    session_time TIME NOT NULL,
    client_id INT REFERENCES Clients(id),
    employee_id INT REFERENCES Employees(id),
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(50),
    comments TEXT,
    attached_files TEXT -- Пути к файлам
);

-- Создание таблицы медицинских карт
CREATE TABLE medical_cards (
    id SERIAL PRIMARY KEY,
    client_id INT REFERENCES Clients(id),
    health_info TEXT NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы записей медицинских карт
-- CREATE TABLE Medical_Card_Records (
--     id SERIAL PRIMARY KEY,
--     card_id INT REFERENCES Medical_Cards(card_id),
--     session_id INT REFERENCES Medical_Sessions(session_id),
--     record_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     details TEXT NOT NULL
-- );

-- Создание таблицы смен сотрудников
CREATE TABLE shifts (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES Employees(id),
    shift_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
);

-- Создание таблицы выходных, отпусков и больничных сотрудников
CREATE TABLE time_off (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES Employees(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    type VARCHAR(50) NOT NULL -- Тип (отпуск, больничный и т.д.)
);

