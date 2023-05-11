# Roba no hashi

From the German word "Eselsbrücke" meaning mnemonic.

This App allows users to lookup Radicals, Kanji and Vocabulary and mnemonics to aid learning.
The difference from something like WaniKani is that user can upload their own menmonics and share with a community.

The app currently runs on web as an alpa here: [https://robanohashi.com](https://robanohashi.com)
API documentation: [.robanohashi.org/docs/index.html](.robanohashi.org/docs/index.html)

### Redis schema

<?xml version="1.0" encoding="UTF-8"?>
<!-- Do not edit this file with editors other than diagrams.net -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="648px" height="420px" viewBox="-0.5 -0.5 648 420" content="&lt;mxfile host=&quot;app.diagrams.net&quot; modified=&quot;2023-05-11T19:13:05.234Z&quot; agent=&quot;Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.78&quot; etag=&quot;5-17_8YKnV42hCe0Ojtu&quot; version=&quot;21.2.3&quot; type=&quot;google&quot;&gt;&lt;diagram name=&quot;Page-1&quot; id=&quot;8-u5jU3yXvecyc1s5xK5&quot;&gt;7VvbcuI4EP0aHuOSJd94DCSTrd2Z3dRmqmb3iRK2AAXbYmVBYL5+JUsGy064TC4wDDxQ6Fhq2d19uttt00H9bHnH8WzyhSUk7UCQLDvopgNhGCH5rYCVBgIQaGDMaaIhdwM80O/EgMCgc5qQwpooGEsFndlgzPKcxMLCMOfsyZ42Yqm96wyPSQt4iHHaRr/RREw0Gvlgg/9G6HhS7ewCcyTD1WQDFBOcsKcahG47qM8ZE/pXtuyTVOmu0ote9+mFo+sT4yQX+yxIPj3eCeCKwfXffxQjGN2tvn65Co01CrGqrpgkUgFmyLiYsDHLcXq7QXuczfOEKLFAjjZzPjM2k6ArwUcixMpYE88Fk9BEZKk5qvdUG714Le5aQ9KzCMuI4Cs5hZMUC7qw12Fj4/F63nrpPaNSIgTGHT0X6iXGG2EX2CIKNucxMavqymwLcgDyvAiGfhRAEFpiPdgQKzAfE9ESe805XtWmzdSEYv/TRz7YdZbb5ssf+gyqUU3VG6h0o0NcKtJ7LnA6N3aKWTZjubTtgCbFsw73GQ9l3LCcBKd0nMvfsVxHuAQWhAsqmXltDmQ0SbQ/koJ+x8NSnvJIo0Yp3O91/Ju1wykBZPlc2DCLN2Stu+IW4rQd1Ii/Ao7r+b4W9qNOW01ho1FBRMNyb2KrCP5a9A9snrp78vQ5QVvoD33ggNqncZ4vxJiDg0HzYqJgxzlHr5vv7Qg2zfPxPiDYRH4r2OAMp2P5JSjLzyfeaKZuizeytkK2AU49+gQt201x/kg76LoDA5zJYNJLpU56Ukg1HotSu0GJD7n8ZSEjVp56zFLGS4nBf3OmJyD1GY3qkF77+8Nff1br5WVoEZXYhuvYlu5J4/WB46tvZXfYV5gDvRKwwXKSA6MG3O3qA17YlKLxED4n3K3A0tkqPy19+p4VVHm+PMchE4JlckIxJSKeGGeteXzTwYWK3z1czHQhPaJLFeXbxCi1QPjtgmhlKFmyup0pJWXLsboPcGJaxMztOlOyUnahadpfGwUB4IMQqVWCsympjuSyRthGnj0SxLrw8ZsRWY+fNjV8BU1q5TuEjv8yQSz/P7gsCo+RaqW++Ooftd7pVsN/y6HneRVwszQb6NGqPronnEoFKLuX4BFzd3RIxb13Id/IctDvSlLuVRUcnrDtrVC0K6Funf9O1Xv3F0momo9bE2roBpYBruCpZ1QXtazHcaJ0fkmql6T6xkl1HXOPl1Qj0PKmj02qoYvqaVXd9kc/WV51Pcdr3On9cDKVslzUuvGtUivynSioHUV7pdk3u1X0WrHxDPtSEXjeSzZpLQiRndZO/j7Rbd8oLliMh/MUK7mXxHZJbK9LbNBObMHR8xps99AzgnOajwdZTjKW03iQsKd8sGBStUWLA+tJb0OG0UjRoU2Gh9uvcon0tHlBuNpLed6FG+fEDTc6NW4gdzc35rMLMy7M+FhmuN3jU6Nd4pb+N2RsmmE+bdOhcs+PoILpRJTbgU4oLxGU/f46UNZ1F76cI1+C6s2Y0+nJo3B3KhnhBeNUkONyp562LtQ4N2q4UejYyeQEemvBUV4N2/TWOtbjKr/zjk01eZFlh2uLNszLDEA3qXY2K17bpjv02RIK7OCKzA4vdeqa80O34S/v8Wip/XCimA8fJXvLKHoO3bdg51thvg/tJ5dGKafbfQuO/OD6J40DRwkD0LVfAYLBR/Aanj+vtz8sBg4CHrQ0f3XyXfXKk7cUv5fe+qW0ffPSFrnvWNrK4eYfFZopm7+loNv/AQ==&lt;/diagram&gt;&lt;/mxfile&gt;" style="background-color: rgb(255, 255, 255);"><defs><style type="text/css">@import url(https://fonts.googleapis.com/css?family=Architects+Daughter);&#xa;</style></defs><g><path d="M 338 230 L 338.03 290 L 338.03 353.63" fill="none" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 338.03 358.88 L 334.53 351.88 L 338.03 353.63 L 341.53 351.88 Z" fill="rgb(0, 0, 0)" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 286px; margin-left: 338px;"><div data-drawio-colors="color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; background-color: rgb(255, 255, 255); white-space: nowrap;">component_ids</div></div></div></foreignObject><text x="338" y="290" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="11px" text-anchor="middle">component_ids</text></switch></g><path d="M 393.03 190 L 393 126 L 394.03 126 L 394.03 80 L 393 80 L 393 66.37" fill="none" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 393 61.12 L 396.5 68.12 L 393 66.37 L 389.5 68.12 Z" fill="rgb(0, 0, 0)" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 139px; margin-left: 392px;"><div data-drawio-colors="color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; background-color: rgb(255, 255, 255); white-space: nowrap;">amalgamation_ids</div></div></div></foreignObject><text x="392" y="143" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="11px" text-anchor="middle">amalgamation_ids</text></switch></g><rect x="341" y="190" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 352.7 190 C 349.6 190 346.62 191.19 344.43 193.3 C 342.23 195.4 341 198.27 341 201.25 C 341 204.23 342.23 207.1 344.43 209.2 C 346.62 211.31 349.6 212.5 352.7 212.5 C 357.13 212.49 361.18 210.08 363.16 206.26 L 367.58 206.26 L 371.3 203.09 L 374.18 205.55 L 376.98 203.09 L 379.71 205.55 L 382.62 203.09 L 385.51 205.55 L 390.35 201.42 C 390.97 200.85 391 200.39 390.35 199.77 L 386.25 196.32 L 363.2 196.32 C 361.24 192.45 357.17 190 352.7 190 Z M 347.56 198.4 C 349.2 198.4 350.53 199.67 350.53 201.25 C 350.53 202.83 349.2 204.11 347.56 204.11 C 345.92 204.11 344.59 202.83 344.59 201.25 C 344.59 199.67 345.92 198.4 347.56 198.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 220px; margin-left: 366px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">kanji:&lt;id&gt;<br /><font color="#3333ff">JSON</font></div></div></div></foreignObject><text x="366" y="232" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">kanji:&lt;i...</text></switch></g><path d="M 394.03 360 L 394.03 320 L 394 206.36" fill="none" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 394 201.11 L 397.5 208.11 L 394 206.36 L 390.5 208.11 Z" fill="rgb(0, 0, 0)" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 302px; margin-left: 396px;"><div data-drawio-colors="color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; background-color: rgb(255, 255, 255); white-space: nowrap;">amalgamation_ids</div></div></div></foreignObject><text x="396" y="306" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="11px" text-anchor="middle">amalgamation_ids</text></switch></g><rect x="341" y="360" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 352.7 360 C 349.6 360 346.62 361.19 344.43 363.3 C 342.23 365.4 341 368.27 341 371.25 C 341 374.23 342.23 377.1 344.43 379.2 C 346.62 381.31 349.6 382.5 352.7 382.5 C 357.13 382.49 361.18 380.08 363.16 376.26 L 367.58 376.26 L 371.3 373.09 L 374.18 375.55 L 376.98 373.09 L 379.71 375.55 L 382.62 373.09 L 385.51 375.55 L 390.35 371.42 C 390.97 370.85 391 370.39 390.35 369.77 L 386.25 366.32 L 363.2 366.32 C 361.24 362.45 357.17 360 352.7 360 Z M 347.56 368.4 C 349.2 368.4 350.53 369.67 350.53 371.25 C 350.53 372.83 349.2 374.11 347.56 374.11 C 346.78 374.11 346.02 373.81 345.46 373.27 C 344.91 372.74 344.59 372.01 344.59 371.25 C 344.59 369.67 345.92 368.4 347.56 368.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 390px; margin-left: 366px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">radical:&lt;id&gt;<br /><font color="#3333ff">JSON</font></div></div></div></foreignObject><text x="366" y="402" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">radical:...</text></switch></g><path d="M 340.48 60 L 341 118.48 L 340.23 169.49" fill="none" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 340.15 174.74 L 336.75 167.69 L 340.23 169.49 L 343.75 167.8 Z" fill="rgb(0, 0, 0)" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 80px; margin-left: 342px;"><div data-drawio-colors="color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; background-color: rgb(255, 255, 255); white-space: nowrap;">component_ids</div></div></div></foreignObject><text x="342" y="83" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="11px" text-anchor="middle">component_ids</text></switch></g><rect x="351" y="0" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 362.7 0 C 359.6 0 356.62 1.19 354.43 3.3 C 352.23 5.4 351 8.27 351 11.25 C 351 14.23 352.23 17.1 354.43 19.2 C 356.62 21.31 359.6 22.5 362.7 22.5 C 367.13 22.49 371.18 20.08 373.16 16.26 L 377.58 16.26 L 381.3 13.09 L 384.18 15.55 L 386.98 13.09 L 389.71 15.55 L 392.62 13.09 L 395.51 15.55 L 400.35 11.42 C 400.97 10.85 401 10.39 400.35 9.77 L 396.25 6.32 L 373.2 6.32 C 371.24 2.45 367.17 0 362.7 0 Z M 357.56 8.4 C 359.2 8.4 360.53 9.67 360.53 11.25 C 360.53 12.83 359.2 14.11 357.56 14.11 C 355.92 14.11 354.59 12.83 354.59 11.25 C 354.59 9.67 355.92 8.4 357.56 8.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 30px; margin-left: 376px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">vocabulary:&lt;id&gt;<br /><font color="#3333ff">JSON</font></div></div></div></foreignObject><text x="376" y="42" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">vocabula...</text></switch></g><rect x="111" y="0" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 122.7 0 C 119.6 0 116.62 1.19 114.43 3.3 C 112.23 5.4 111 8.27 111 11.25 C 111 14.23 112.23 17.1 114.43 19.2 C 116.62 21.31 119.6 22.5 122.7 22.5 C 127.13 22.49 131.18 20.08 133.16 16.26 L 137.58 16.26 L 141.3 13.09 L 144.18 15.55 L 146.98 13.09 L 149.71 15.55 L 152.62 13.09 L 155.51 15.55 L 160.35 11.42 C 160.97 10.85 161 10.39 160.35 9.77 L 156.25 6.32 L 133.2 6.32 C 131.24 2.45 127.17 0 122.7 0 Z M 117.56 8.4 C 119.2 8.4 120.53 9.67 120.53 11.25 C 120.53 12.83 119.2 14.11 117.56 14.11 C 115.92 14.11 114.59 12.83 114.59 11.25 C 114.59 9.67 115.92 8.4 117.56 8.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 30px; margin-left: 136px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">meaning_mnemonic_down_voters:&lt;mnemonic_id&gt;<br /><font color="#ff3333">SET [user_id]</font></div></div></div></foreignObject><text x="136" y="42" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">meaning_...</text></switch></g><rect x="111" y="130" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 122.7 130 C 119.6 130 116.62 131.19 114.43 133.3 C 112.23 135.4 111 138.27 111 141.25 C 111 144.23 112.23 147.1 114.43 149.2 C 116.62 151.31 119.6 152.5 122.7 152.5 C 127.13 152.49 131.18 150.08 133.16 146.26 L 137.58 146.26 L 141.3 143.09 L 144.18 145.55 L 146.98 143.09 L 149.71 145.55 L 152.62 143.09 L 155.51 145.55 L 160.35 141.42 C 160.97 140.85 161 140.39 160.35 139.77 L 156.25 136.32 L 133.2 136.32 C 131.24 132.45 127.17 130 122.7 130 Z M 117.56 138.4 C 119.2 138.4 120.53 139.67 120.53 141.25 C 120.53 142.83 119.2 144.11 117.56 144.11 C 115.92 144.11 114.59 142.83 114.59 141.25 C 114.59 139.67 115.92 138.4 117.56 138.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 160px; margin-left: 136px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">meaning_mnemonic_up_voters:&lt;mnemonic_id&gt;<br /><font color="#ff3333">SET [user_id]</font></div></div></div></foreignObject><text x="136" y="172" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">meaning_...</text></switch></g><rect x="526" y="190" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 537.7 190 C 534.6 190 531.62 191.19 529.43 193.3 C 527.23 195.4 526 198.27 526 201.25 C 526 204.23 527.23 207.1 529.43 209.2 C 531.62 211.31 534.6 212.5 537.7 212.5 C 542.13 212.49 546.18 210.08 548.16 206.26 L 552.58 206.26 L 556.3 203.09 L 559.18 205.55 L 561.98 203.09 L 564.71 205.55 L 567.62 203.09 L 570.51 205.55 L 575.35 201.42 C 575.97 200.85 576 200.39 575.35 199.77 L 571.25 196.32 L 548.2 196.32 C 546.24 192.45 542.17 190 537.7 190 Z M 532.56 198.4 C 534.2 198.4 535.53 199.67 535.53 201.25 C 535.53 202.83 534.2 204.11 532.56 204.11 C 530.92 204.11 529.59 202.83 529.59 201.25 C 529.59 199.67 530.92 198.4 532.56 198.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 220px; margin-left: 551px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">user_bookmarks:&lt;user_id&gt;<br /><font color="#ff3333">SET [radical_id | kanji_id | vocab_id]</font></div></div></div></foreignObject><text x="551" y="232" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">user_boo...</text></switch></g><rect x="113.5" y="360" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 125.2 360 C 122.1 360 119.12 361.19 116.93 363.3 C 114.73 365.4 113.5 368.27 113.5 371.25 C 113.5 374.23 114.73 377.1 116.93 379.2 C 119.12 381.31 122.1 382.5 125.2 382.5 C 129.63 382.49 133.68 380.08 135.66 376.26 L 140.08 376.26 L 143.8 373.09 L 146.68 375.55 L 149.48 373.09 L 152.21 375.55 L 155.12 373.09 L 158.01 375.55 L 162.85 371.42 C 163.47 370.85 163.5 370.39 162.85 369.77 L 158.75 366.32 L 135.7 366.32 C 133.74 362.45 129.67 360 125.2 360 Z M 120.06 368.4 C 121.7 368.4 123.03 369.67 123.03 371.25 C 123.03 372.83 121.7 374.11 120.06 374.11 C 118.42 374.11 117.09 372.83 117.09 371.25 C 117.09 369.67 118.42 368.4 120.06 368.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 390px; margin-left: 139px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">meaning_mnemonic_favorites:&lt;user_id&gt;<br /><font color="#ff3333">SET [mnemonic_id]</font></div></div></div></foreignObject><text x="139" y="402" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">meaning_...</text></switch></g><path d="M 163.5 261.24 L 286.03 261.24 L 286.03 11.31 L 344.63 11.26" fill="none" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 349.88 11.25 L 342.89 14.76 L 344.63 11.26 L 342.88 7.76 Z" fill="rgb(0, 0, 0)" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 259px; margin-left: 262px;"><div data-drawio-colors="color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; background-color: rgb(255, 255, 255); white-space: nowrap;">subject_id</div></div></div></foreignObject><text x="262" y="262" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="11px" text-anchor="middle">subject_id</text></switch></g><path d="M 138.52 250 L 138.52 201.31 L 334.63 201.25" fill="none" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="stroke"/><path d="M 339.88 201.25 L 332.88 204.75 L 334.63 201.25 L 332.88 197.75 Z" fill="rgb(0, 0, 0)" stroke="rgb(0, 0, 0)" stroke-miterlimit="10" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe center; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 203px; margin-left: 254px;"><div data-drawio-colors="color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 11px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; background-color: rgb(255, 255, 255); white-space: nowrap;">subject_id</div></div></div></foreignObject><text x="254" y="206" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="11px" text-anchor="middle">subject_id</text></switch></g><rect x="113.5" y="250" width="50" height="22.5" fill="none" stroke="none" pointer-events="all"/><path d="M 125.2 250 C 122.1 250 119.12 251.19 116.93 253.3 C 114.73 255.4 113.5 258.27 113.5 261.25 C 113.5 264.23 114.73 267.1 116.93 269.2 C 119.12 271.31 122.1 272.5 125.2 272.5 C 129.63 272.49 133.68 270.08 135.66 266.26 L 140.08 266.26 L 143.8 263.09 L 146.68 265.55 L 149.48 263.09 L 152.21 265.55 L 155.12 263.09 L 158.01 265.55 L 162.85 261.42 C 163.47 260.85 163.5 260.39 162.85 259.77 L 158.75 256.32 L 135.7 256.32 C 133.74 252.45 129.67 250 125.2 250 Z M 120.06 258.4 C 121.7 258.4 123.03 259.67 123.03 261.25 C 123.03 262.83 121.7 264.11 120.06 264.11 C 118.42 264.11 117.09 262.83 117.09 261.25 C 117.09 259.67 118.42 258.4 120.06 258.4 Z" fill="#005073" stroke="none" pointer-events="all"/><g transform="translate(-0.5 -0.5)"><switch><foreignObject pointer-events="none" width="100%" height="100%" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility" style="overflow: visible; text-align: left;"><div xmlns="http://www.w3.org/1999/xhtml" style="display: flex; align-items: unsafe flex-start; justify-content: unsafe center; width: 1px; height: 1px; padding-top: 280px; margin-left: 139px;"><div data-drawio-colors="color: rgb(0, 0, 0); " style="box-sizing: border-box; font-size: 0px; text-align: center;"><div style="display: inline-block; font-size: 12px; font-family: Helvetica; color: rgb(0, 0, 0); line-height: 1.2; pointer-events: all; white-space: nowrap;">meaning_mnemonic:&lt;id&gt;<br /><font color="#3333ff">JSON</font></div></div></div></foreignObject><text x="139" y="292" fill="rgb(0, 0, 0)" font-family="Helvetica" font-size="12px" text-anchor="middle">meaning_...</text></switch></g></g><switch><g requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"/><a transform="translate(0,-5)" xlink:href="https://www.diagrams.net/doc/faq/svg-export-text-problems" target="_blank"><text text-anchor="middle" font-size="10px" x="50%" y="100%">Text is not SVG - cannot display</text></a></switch></svg>

### Authentication
Authentication is handled entirely by Firebase. This api only validates the JSON web token for endpoints that require it.

### Development

For development I suggest to start a redis docker container.
```bash
docker run -d \
  --name redis \
  -p 6379:6379 \
  -p 8001:8001 \
  -v "$(pwd)/data:/data" \
  redis/redis-stack:latest
```

#### Migrations
Migrations for the Redis instance are performed locally and the `dump.rdb` files are backed on `S3`.
Create a new package in `/cmd/<package>` for the migration code.

#### Docs
1. Download the `swag` cli.
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
2. Run `swag init` to generate the files

#### Indices
Run `go run ./cmd/indices` to recreate indices

#### Deployment
```bash
docker build -t martingrzzler/robanohashi-api:latest .
```
```bash
docker push martingrzzler/robanohashi-api:latest
```

On the production server update individual docker swarm services:
```bash
docker service update --image martingrzzler/robanohashi-api:latest robanohashi_api
```

```bash
docker service update robanohashi_redis --force
```

#### Test
Integration tests:

Make sure to start an empty redis instance.
```bash
go test ./persist
```

Unit tests:
```bash
go test ./internal/utils
```

