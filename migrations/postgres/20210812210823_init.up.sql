CREATE TABLE budgets (
    codename varchar(255) PRIMARY KEY,
    daily_limit integer NOT NULL
);

CREATE TABLE categories (
    codename varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    is_base_expense boolean NOT NULL,
    aliases text NOT NULL
);

CREATE TABLE expenses (
    id serial PRIMARY KEY,
    amount integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    category_codename varchar(255) NOT NULL,
    raw_text text NOT NULL,
    FOREIGN KEY (category_codename) REFERENCES categories(codename)
);

INSERT INTO categories (codename, name, is_base_expense, aliases)
VALUES
    ('products', 'продукты', true, 'еда'),
    ('coffee', 'кофе', true, ''),
    ('dinner', 'обед', true, 'столовая, ланч, бизнес-ланч, бизнес ланч'),
    ('cafe', 'кафе', true, 'ресторан, рест, мак, макдональдс, макдак, kfc, ilpatio, il patio'),
    ('transport', 'общ. транспорт', false, 'метро, автобус, metro'),
    ('taxi', 'такси', false, 'яндекс такси, yandex taxi'),
    ('phone', 'телефон', false, 'теле2, связь'),
    ('books', 'книги', false, 'литература, литра, лит-ра'),
    ('internet', 'интернет', false, 'инет, inet'),
    ('subscriptions', 'подписки', false, 'подписка'),
    ('other', 'прочее', true, '');

INSERT INTO budgets (codename, daily_limit) VALUES ('base', 500);