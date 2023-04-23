import datasets
import json

_DESCRIPTION = """\
Contains Kanji images with corresponding radicals ids from WaniKani or https://api.robanohashi.org/docs/index.html
"""

_METADATA_URL = "https://huggingface.co/datasets/martingrzzler/kanjis2radicals/raw/main/kanji_metadata.jsonl"
_IMAGES_URL = "https://huggingface.co/datasets/martingrzzler/kanjis2radicals/resolve/main/kanjis.tar.gz"


class Kanji2Radicals(datasets.GeneratorBasedBuilder):
    """Kanji to radicals dataset."""

    def _info(self):
        return datasets.DatasetInfo(
            description=_DESCRIPTION,
            features=datasets.Features(
                {
                    "kanji_image": datasets.Image(),
                    "meta": {
                        "id": datasets.Value("int32"),
                        "characters": datasets.Value("string"),
                        "meanings": datasets.Value("string"),
                        "radicals": datasets.Sequence(
                            {
                                "characters": datasets.Value("string"),
                                "id": datasets.Value("int32"),
                                "slug": datasets.Value("string"),
                            }
                        ),
                    },
                }
            ),
            supervised_keys=None,
            homepage="https://robanohashi.org/",
        )

    def _split_generators(self, dl_manager):
        metadata_path = dl_manager.download(_METADATA_URL)
        images_path = dl_manager.download(_IMAGES_URL)
        images_iter = dl_manager.iter_archive(images_path)

        return [
            datasets.SplitGenerator(
                name=datasets.Split.TRAIN,
                gen_kwargs={
                    "metadata_path": metadata_path,
                    "images_iter": images_iter,
                },
            ),
        ]

    def _generate_examples(self, metadata_path, images_iter):
        meta_kanjis = {}

        with open(metadata_path, encoding="utf-8") as f:
            for line in f:
                metadata = json.loads(line)
                meta_kanjis[metadata["characters"]] = metadata

        for idx, (image_path, image) in enumerate(images_iter):
            characters = image_path.split("/")[-1].split(".")[0]
            yield image_path, {
                "meta": meta_kanjis[characters],
                "kanji_image": image.read(),
            }
