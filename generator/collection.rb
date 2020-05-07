require 'erb'

class CollectionGenerator
  Param = Struct.new(:method_name, :variable_name, :param_type, :param_name)

  def initialize(name:, plural_name: nil, path:, var_prefix: nil, api_url: nil, entity:, args: {}, params: {})
    @name = name
    @plural_name = plural_name
    @path = path
    @var_prefix = var_prefix
    @entity = entity
    @args = args
    @api_url = api_url
    @params = params
  end

  def collection_api_url
    @api_url
  end

  def collection_var_prefix
    @var_prefix
  end

  def collection_type_name
    @name
  end

  def collection_type_name_plural
    if @plural_name.nil?
      "#{@name}s"
    else
       @plural_name
    end
  end

  def collection_element_type
    @entity
  end

  def collection_element_uncapitalized
    @collection_element_uncapitalized ||= uncapitalize(@entity)
  end

  def resource_path
    @path
  end

  def collection_args_prepared
    @collection_args_prepared ||= @args.map { |a, t| "#{a} #{t}" }
  end

  def collection_args
    @args
  end

  def collection_params
    @collection_params ||= @params.each_with_object([]) do |(name, param_type), acc|
      acc << Param.new(
        camelize(name.to_s),
        camelize(name == :type ? 't' : name.to_s, false),
        param_type,
        name.to_s
      )
    end
  end

  def render
    ERB.new(File.read(File.join(File.dirname(__FILE__), 'templates', 'collection.erb')), nil, '-').result(binding)
  end

  def render_to_file(file)
    File.write(file, render)
  end

  private

  def uncapitalize(string)
    "#{string[0, 1].downcase}#{string[1..-1]}"
  end

  def camelize(string, uppercase_first_letter = true)
    if uppercase_first_letter
      string = string.sub(/^[a-z\d]*/) { |match| match.capitalize }
    else
      string = string.sub(/^(?:(?=\b|[A-Z_])|\w)/) { |match| match.downcase }
    end
    string.gsub(/(?:_|(\/))([a-z\d]*)/) { "#{$1}#{$2.capitalize}" }.gsub("/", "::")
  end
end
