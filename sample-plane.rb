require "ipaddr"

def search_ip(ip)
  open("./ips-rb", "r") do |file|
    while line = file.gets
      return true if IPAddr.new(line).include?(ip)
    end
  end

  return false
end

