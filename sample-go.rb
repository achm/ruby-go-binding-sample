require "ffi"

module GeoIP
  extend FFI::Library
  ffi_lib "./libips.so"
  attach_function :search, [:string], :bool
end

