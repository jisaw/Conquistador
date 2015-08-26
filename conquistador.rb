require "sqlite3"

class Conquistador
    def initialize(db)
        @db = db
    end

    def init_tables(db)
        begin
            @db.execute "create table if not exists Goal(id INTEGER PRIMARY KEY, goal TEXT, goal_date INTEGER, complete INTEGER DEFAULT 0, user_id INTEGER)"
            @db.execute "create table if not exists User(id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT, username TEXT, password TEXT, created INTEGER)"
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

    def complete_goal(id)
        @db.execute "update goal set complete = 1 where id = #{id}"
    end

    def add_user(first_name, last_name, email, username, password)
        created = Time.now.to_i
        @db.execute "insert into User(first_name, last_name, email, username, password, created) values(#{first_name}, #{last_name}, #{email}, #{username}, #{password}, #{created})"
    end

    def delete_user(id)
        @db.execute "delete from User where id = #{id}"
    end

    def update_password(id, password)
        @db.execute "update user set password = #{password} where id = #{id}"
    end

    def update_email(id, email)
        @db.execute "update user set email = #{email} where id = #{id}"
    end

    def update_username(id, username)
        @db.execute "update user set username = #{username} where id = #{id}"
    end

    def login(username, password)
        stm = @db.prepare "select password from user where username = #{username}"
        rs = stm.execute

        rs.do |pass|
            if password = pass
                result = true
                break
            end
        if result
            return true
        else
            return false
        end
    end

end
