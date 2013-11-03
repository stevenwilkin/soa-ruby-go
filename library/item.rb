class Item < Struct.new(:id, :text)
  def to_s
    "#{id} - #{text}"
  end
end
