const assembler = (commands) => {
  let register = {};
  const splitCommands = commands.map((command) => command.split(" "));

  for (let i = 0; i < splitCommands.length; i++) {
    let cmd = splitCommands[i];
    switch (cmd[0]) {
      case "mov":
        if (cmd[2] in register) {
          register[cmd[1]] = register[cmd[2]];
        } else {
          register[cmd[1]] = parseInt(cmd[2]);
        }
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
        if (regVal !== 0 && vl < 0) {
          i += vl - 1;
        } else if (regVal !== 0 && vl > 0) {
          i += vl - 1;
        }
        break;
    }

    // console.log(i, cmd[0], cmd[1], cmd[2], register);
  }
  return register;
};
