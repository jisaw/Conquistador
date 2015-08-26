require "sqlite3"

class Dao
    
    def initialize(dbname)
        @dbname = dbname
        begin 
            @db = SQLite3::Database.open @dbname
            @db.execute "create table if not exists Goal(id INTEGER PRIMARY KEY, goal TEXT, goal_date INTEGER, complete INTEGER DEFAULT 0, user_id INTEGER)"
            @db.execute "create table if not exists User(id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT, password TEXT)"
    end



end
