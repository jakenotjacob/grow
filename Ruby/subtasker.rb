
class TaskRetry > StandardError; end
class TaskFailed > StandardError
  def message
    "Task failed to execute"
  end
end

module Task
  module Chef
    def run
      delete_node
      delete_client
    end
    def self.delete_node
    end
    def self.delete_client
    end
  end

  module Sensu
  end
end

module Task
  # Task list
  def self.list
    self.constants
  end

  # This is the entrypoint where each Task's run() method is called:
  #
  # If a Task raises a TaskRetry inside the run() method, it
  # will attempt to re-run all subtasks up to 3 times.
  #
  # Call each Task::$task's run() method, where it places
  # subtasks taken to complete the full task
  def self.run_all
    Task.list.each do |task|
      retries = 0
      begin
        eval "Task::#{task}.run"
      rescue TaskRetry
        retry if retries < 2
        retries += 1
      # Send all other errors to log (ignore failure)
      rescue => e
        @out_file.puts "Task failed: #{e.inspect}"
      else
        # Handle success
        #TODO Slack.postmessagethinger
      ensure
        # Cleanup after running tasks
      end
      # Return value
    end
    return true
  end

  module Chef
    def self.run
      delete_node
      delete_client
    rescue
      raise TaskRetry
    ensure
      return true
    end

    def delete_node
    end

    def delete_client
    end

  end

  #module Sensu
  #  def self.run
  #    ... do work
  #  rescue
  #    ... handle error
  #  else
  #    ... handle success
  #  ensure
  #    ... cleanup
  #  end
  #
  #  def silence
  #  end
  #end
end





class Decom
  def initialize()
    raise AccessDeniedError, "unauthorized for use of this job" \
      unless user.admin?
  end
end


module Task

  # Iterate over namespace for Tasks
  def run_all
    self.constants.each do |task|

    end
  end

  module Chef
    SUBTASKS = %w(delete_node)
    def delete_node
      "dalat that"
    end
  end

  module Sensu
    SUBTASKS = %w(silence)
    def silence
      "silencio"
    end
  end
end

class Decom
  Task.run_all
end



class Decom
  include Task::Chef
  def run
    retries ||= 0
    #TODO Guard this by ensuring the node is turned off?
    Task::Chef.delete_node!
    Task::Chef.delete_client!
  rescue TaskRetry
    # TODO Tasks wanting to enter retry loop just need to "raise"
    # Only OK for idempotent methods; lets not re-run disk wipe 3 times
    retry if retries < 2
    retries += 1
  else
    # Notify Slack of Task success
    # If no error, save as success
  ensure
    # cleanup
  end

end

Decom::Tasks.each do |task|
  retries ||= 0
  begin
    task.run
  rescue
    if retries < 2
      retry
    else
      # Slack notify failure, move to next task
      next
    end
  ensure
  end
end

