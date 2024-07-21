import sqlite3

# データベース接続の作成
conn = sqlite3.connect('photos.db')
cursor = conn.cursor()

# Photosテーブル作成
create_photos_table_query = '''
CREATE TABLE photos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    filepath TEXT NOT NULL,
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

# コミットして接続を閉じる
conn.commit()
conn.close()
