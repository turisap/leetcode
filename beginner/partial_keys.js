// You need to make a function that takes an object as an argument, and returns a very similar object but with a special property. The returned object should allow a user to access values by providing only the beginning of the key for // the value they want. For example if the given object has a key idNumber, you should be able to access its value on the returned object by using a key idNum or even simply id. Num and Number shouldn't work because we are only looki// ng for matches at the beginning of a key.
const partial_keys = (obj) => {
  return new Proxy(obj, {
    get(target, name) {
      const match = Object.entries(target)
        .sort((a, b) => a[0].localeCompare(b[0]))

        .find(([prop, value]) => prop.startsWith(name));

      return match ? match[1] : null;
    },
  });
};

module.exports = partial_keys;
