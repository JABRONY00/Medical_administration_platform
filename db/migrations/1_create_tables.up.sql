-- Создание таблицы клиентов
CREATE TABLE clients (
    id uuid DEFAULT gen_random_uuid(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(100),

    birth_date DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    address TEXT,

    PRIMARY KEY (id)
);

-- Создание таблицы отделов
CREATE TABLE departments (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    description TEXT,

    PRIMARY KEY (id)
);

-- Создание таблицы ролей пользователей
CREATE TABLE roles (
    id SERIAL,
    role_name VARCHAR(50) NOT NULL,

    PRIMARY KEY (id)
);

-- Создание таблицы сотрудников
CREATE TABLE employees (
    id uuid DEFAULT gen_random_uuid(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(15),
    password_hash TEXT NOT NULL,

    birth_date DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    position VARCHAR(100),

    department_id INT,
    role_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Создание таблицы кабинетов
CREATE TABLE rooms (
    id SERIAL,
    room_number VARCHAR(10) NOT NULL,

    department_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (department_id) REFERENCES departments(id)
);

-- Связующая таблица между сотрудниками и кабинетами (многие ко многим)
CREATE TABLE employee_rooms (
    id SERIAL,
    employee_id uuid,
    room_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

-- Создание таблицы расходных материалов в кабинетах
CREATE TABLE room_storage (
    id SERIAL,
    item_name VARCHAR(100) NOT NULL,
    quantity INT NOT NULL,
    room_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Создание таблицы заявок на посещение врача
CREATE TABLE appointments (
    id SERIAL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(50),

    client_id uuid,
    employee_id uuid,
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

-- Создание таблицы посещений врача
CREATE TABLE medical_sessions (
    id SERIAL,
    session_date DATE NOT NULL,
    session_time TIME NOT NULL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(50),
    comments TEXT,
    attached_files TEXT, -- Пути к файлам

    client_id uuid,
    employee_id uuid,
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

-- Создание таблицы медицинских карт
CREATE TABLE medical_cards (
    id SERIAL,
    health_info TEXT NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    client_id uuid,
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

-- Создание таблицы смен сотрудников
CREATE TABLE shifts (
    id SERIAL,
    shift_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,

    employee_id uuid,
    PRIMARY KEY (id),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Создание таблицы выходных, отпусков и больничных сотрудников
CREATE TABLE time_off (
    id SERIAL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    type VARCHAR(50) NOT NULL, -- Тип (отпуск, больничный и т.д.)

    employee_id uuid,
    PRIMARY KEY (id),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE ON UPDATE CASCADE
);

