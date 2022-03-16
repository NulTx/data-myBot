<<__EntryPoint>>
function main(): noreturn {
 $user_one = shape('name' => "Joe", 'age' => 25);
 print("name: " . $user_one['name'] . "\n");
 print("age " . $user_one['age'] . "\n");
 exit(0);
}
