
fn ->(Hello)[name string, age int] : string {
  print("Hello"+name+" ton âge est: "+age)  
}

fn ->(main)[] : void {
  print("Hello world")
  hello("Bryton", 18)
}
