const assembler = (commands) => {
  let register = {};
  const splitCommands = commands.map((command) => command.split(" "));

  for (let i = 0; i < splitCommands.length; i++) {
    let cmd = splitCommands[i];
    switch (cmd[0]) {
      case "mov":
        register[cmd[1]] = cmd[2];
        break;
      case "inc":
        register[cmd[1]]++;
        break;
      case "dec":
        register[cmd[1]]--;
        break;
      case "jnz":
        const vl = parseInt(cmd[2]);
        const regVal = register[cmd[1]];
        if (regVal && vl < 0) {
          i += vl - 1;
        } else if (regVal && vl > 0) {
          i++;
        }
        break;
    }

    // console.log(i, cmd[0], cmd[1], cmd[2], register);
  }
  return register;
};

const input = ["mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"];

console.log(assembler(input));
