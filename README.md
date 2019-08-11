# CLI utility for showing mp4 file atoms
To run utility, simply run:
```
$ go run mp4parser.go <Path to mp4 file>
```
Output example, which lists mp4 file atoms:
```
Atom type: ftyp size: 32
Atom type: free size: 8
Atom type: mdat size: 58408514
Atom type: moov size: 549855
        Atom type: mvhd size: 108
        Atom type: trak size: 154090
                Atom type: tkhd size: 92
                Atom type: edts size: 36
                        Atom type: elst size: 28
                Atom type: mdia size: 153954
                        Atom type: mdhd size: 32
                        Atom type: hdlr size: 45
                        Atom type: minf size: 153869
                                Atom type: vmhd size: 20
                                Atom type: dinf size: 36
                                        Atom type: dref size: 28
                                Atom type: stbl size: 153805
                                        Atom type: stsd size: 149
                                        Atom type: stts size: 24
                                        Atom type: stss size: 1288
                                        Atom type: stsc size: 40
                                        Atom type: stsz size: 76176
                                        Atom type: stco size: 76120
        Atom type: trak size: 395255
                Atom type: tkhd size: 92
                Atom type: edts size: 36
                        Atom type: elst size: 28
                Atom type: mdia size: 395119
                        Atom type: mdhd size: 32
                        Atom type: hdlr size: 45
                        Atom type: minf size: 395034
                                Atom type: smhd size: 16
                                Atom type: dinf size: 36
                                        Atom type: dref size: 28
                                Atom type: stbl size: 394974
                                        Atom type: stsd size: 106
                                        Atom type: stts size: 32
                                        Atom type: stsc size: 199780
                                        Atom type: stsz size: 118928
                                        Atom type: stco size: 76120
                                        Atom type: udta size: 394
                                                Atom type: meta size: 386
                                Atom type: udta size: 394
                                        Atom type: meta size: 386
                        Atom type: udta size: 394
                                Atom type: meta size: 386
                Atom type: udta size: 394
                        Atom type: meta size: 386
        Atom type: udta size: 394
                Atom type: meta size: 386
All atoms succefully readed. EOF
```