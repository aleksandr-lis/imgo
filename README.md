# Just for fun

To download an image you can just enter its url.

```bash
git@github.com:lis-space/imgo.git
cd imgo
go build
```

## Instagram

I will download the main image.
If there's a video on this page, I will download it as well.
If there's a carousel, I will download the title image only.

```bash
./imgo https://www.instagram.com/p/B09PvydFHd4/
```

## Pinterest

I will download the main image. But only if this image is public.

```bash
./imgo https://www.pinterest.ru/pin/761460249473825644/
```

## Facebook

I will download the main image. But only if this image is public.

```bash
./imgo "https://www.facebook.com/photo.php?fbid=2325054541158452&set=gm.2596952480526017&type=3&theater"
```

## License

MIT.