### Regex ###
Regex.run ~r{[aeiou]}, "caterpillar"

Regex.scan ~r{[aeiou]}, "caterpillar"

Regex.split ~r{[aeiou]}, "caterpillar"

Regex.replace ~r{[aeiou]}, "caterpillar", "*"


### Atoms ###
:fred
:fred?
:fred!
:"f r e d"
:fr@2
:fr@ed


### Tuples ###
#~3, 4+ and use Maps instead
{status, count, action} = {:ok, 42, "next"}

{status, file} = File.open("somefile.txt")
#Writing matches assuming success
{ :ok, file } = File.open("non-existant-file")


### Lists and Keyword Lists ###
[ name: "Bob", nickname: "Raptorman" ] # Converted to list of two-value tuples
[ {name: "Bob"}, {nickname: "Raptorman"}]


### Maps ###
responses = %{ { :error, :enoent } => :fatal, { :error, :busy } => :retry }
colors = %{ :red => 0xff0000, :green => 0x00ff00, :blue => 0x0000ff }

%{ "one" => 1, :two => 2, {1,1,1} => 3}

person = %{ name: "Joe" }
person[:name]
person.name


### Binaries ###
<< 1,2 >>
byte_size bin #2

bin = <<3 :: size(2), 5 :: size(4), 1 :: size(1) >>
byte_size bin #1
#Binary Join
binary1 <> binary2

#number < atom < reference < function < port < pid < tuple < map < list < binary

### Syntax stuffs ###
#
#number < atom < reference < function < port < pid < tuple < map < list < binary
#
[1, fred: 1, dave: 2]
## = [1, {fred: 1, dave: 2}]
{1, fred: 1, dave: 2}
## = {1, [fred:1, dave: 2]}
#



