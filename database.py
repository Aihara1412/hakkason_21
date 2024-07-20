import sqlite3

# データベース接続の作成
conn = sqlite3.connect('photos.db')
cursor = conn.cursor()

# Usersテーブル作成
create_users_table_query = '''
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
'''
cursor.execute(create_users_table_query)

# Photosテーブル作成
create_photos_table_query = '''
CREATE TABLE photos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    filepath TEXT NOT NULL,
    tags TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    category TEXT,
    uploaded_by INTEGER,
    location TEXT,
    camera_settings TEXT,
    is_public BOOLEAN DEFAULT 1,
    FOREIGN KEY (uploaded_by) REFERENCES users(id)
);
'''
cursor.execute(create_photos_table_query)

# Categoriesテーブル作成（オプション）
create_categories_table_query = '''
CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);
'''
cursor.execute(create_categories_table_query)

# コミットして接続を閉じる
conn.commit()
conn.close()
