require "benchmark"
require "./sample-plane"
require "./sample-go"

NUM = 100

def random_ip
  "#{Random.rand(255)}.#{Random.rand(255)}.#{Random.rand(255)}.#{Random.rand(255)}"
end

Benchmark.bm 5 do |r|
  r.report "Ruby" do
    NUM.times do
      search_ip random_ip
    end
  end

  r.report "Go" do
    NUM.times do
      GeoIP.search random_ip
    end
  end
end
