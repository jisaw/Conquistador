require "sqlite3"

class Conquistador
    def initialize(db)
        @db = db
    end

    def init_tables(db)
        begin
            @db.execute "create table if not exists Goal(id INTEGER PRIMARY KEY, goal TEXT, goal_date INTEGER, complete INTEGER DEFAULT 0, user_id INTEGER)"
            @db.execute "create table if not exists User(id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT, password TEXT)"
        rescue SQLite3::Exception => e
            puts "Exception Occured"
            puts e
        ensure
            db.close if db
        end
    end

    def add_goal(goal, user_id)
        goal_date = Time.now.to_i
        @db.execute "insert into Goal(goal, goal_date, user_id) VALUES(#{goal}, #{goal_date}, #{user_id})"
    end

    def delete_goal(id)
        @db.execute "delete from Goal where id = #{id}"
    end

end
