import lib from "kuroshiro";
const Kuroshiro = lib.default;
import KuromojiAnalyzer from "kuroshiro-analyzer-kuromoji";

import fs from "fs";
import readline from "readline";

async function main() {
  const kuroshiro = new Kuroshiro();
  await kuroshiro.init(new KuromojiAnalyzer());

  const readInterface = readline.createInterface({
    input: fs.createReadStream("subjects.json"),
  });

  const outputStream = fs.createWriteStream("temp");

  readInterface.on("line", async function (line) {
    const subject = JSON.parse(line);

    if (subject.object !== "vocabulary") {
      outputStream.write(JSON.stringify(subject) + "\n");
      return;
    }

    for (let i = 0; i < subject.data.context_sentences.length; i++) {
      subject.data.context_sentences[i].hiragana = await kuroshiro.convert(
        subject.data.context_sentences[i].ja,
        {
          to: "hiragana",
        }
      );
    }

    outputStream.write(JSON.stringify(subject) + "\n");
  });

  readInterface.on("close", function () {
    fs.renameSync("temp", "subjects.json");
  });
}

main();
