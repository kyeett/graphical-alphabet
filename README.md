# Graphical Alphabet

A collection of image operations, implemented as an exercises in 2d graphics.

**Inspired by**:

- http://ssp.impulsetrain.com/porterduff.html
- http://507movements.com/

### Porter/Duff composition

|                               &nbsp;                                |                                     &nbsp;                                      |                                     &nbsp;                                      |                                 &nbsp;                                  |                                   &nbsp;                                    |
| :-----------------------------------------------------------------: | :-----------------------------------------------------------------------------: | :-----------------------------------------------------------------------------: | :---------------------------------------------------------------------: | :-------------------------------------------------------------------------: |
|     ![example:Src](examples/Src.png)<br>[Src](examples/Src.png)     |         ![example:Atop](examples/Atop.png)<br>[Atop](examples/Atop.png)         |         ![example:Over](examples/Over.png)<br>[Over](examples/Over.png)         |         ![example:In](examples/In.png)<br>[In](examples/In.png)         |         ![example:Out](examples/Out.png)<br>[Out](examples/Out.png)         |
|   ![example:Dest](examples/Dest.png)<br>[Dest](examples/Dest.png)   | ![example:DestAtop](examples/DestAtop.png)<br>[DestAtop](examples/DestAtop.png) | ![example:DestOver](examples/DestOver.png)<br>[DestOver](examples/DestOver.png) | ![example:DestIn](examples/DestIn.png)<br>[DestIn](examples/DestIn.png) | ![example:DestOut](examples/DestOut.png)<br>[DestOut](examples/DestOut.png) |
| ![example:Clear](examples/Clear.png)<br>[Clear](examples/Clear.png) |           ![example:Xor](examples/Xor.png)<br>[Xor](examples/Xor.png)           |   ![example:Lighter](examples/Lighter.png)<br>[Lighter](examples/Lighter.png)   |

### Transformation

|                                                 &nbsp;                                                  |                                             &nbsp;                                              |                                               &nbsp;                                                |                                               &nbsp;                                                |
| :-----------------------------------------------------------------------------------------------------: | :---------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------: |
|                 ![example:Normal](examples/Normal.png)<br>[Normal](examples/Normal.png)                 | ![example:RotateDest90](examples/RotateDest90.png)<br>[RotateDest90](examples/RotateDest90.png) | ![example:RotateDest180](examples/RotateDest180.png)<br>[RotateDest180](examples/RotateDest180.png) | ![example:RotateDest270](examples/RotateDest270.png)<br>[RotateDest270](examples/RotateDest270.png) |
|                 ![example:Normal](examples/Normal.png)<br>[Normal](examples/Normal.png)                 |         ![example:Rotate90](examples/Rotate90.png)<br>[Rotate90](examples/Rotate90.png)         |         ![example:Rotate180](examples/Rotate180.png)<br>[Rotate180](examples/Rotate180.png)         |         ![example:Rotate270](examples/Rotate270.png)<br>[Rotate270](examples/Rotate270.png)         |
| ![example:FlipHorizontal](examples/FlipHorizontal.png)<br>[FlipHorizontal](examples/FlipHorizontal.png) | ![example:FlipVertical](examples/FlipVertical.png)<br>[FlipVertical](examples/FlipVertical.png) |

### Resize

|                                             &nbsp;                                              |                                 &nbsp;                                  |                                               &nbsp;                                                |                                                           &nbsp;                                                            |
| :---------------------------------------------------------------------------------------------: | :---------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------------: |
|             ![example:Normal](examples/Normal.png)<br>[Normal](examples/Normal.png)             | ![example:Resize](examples/Resize.png)<br>[Resize](examples/Resize.png) | ![example:ResizeInPlace](examples/ResizeInPlace.png)<br>[ResizeInPlace](examples/ResizeInPlace.png) |               ![example:ResizeOffset](examples/ResizeOffset.png)<br>[ResizeOffset](examples/ResizeOffset.png)               |
| ![example:ResizeLarger](examples/ResizeLarger.png)<br>[ResizeLarger](examples/ResizeLarger.png) | ![example:Padded](examples/Padded.png)<br>[Padded](examples/Padded.png) | ![example:PaddedResized](examples/PaddedResized.png)<br>[PaddedResized](examples/PaddedResized.png) | ![example:ResizeLargerCorrect](examples/ResizeLargerCorrect.png)<br>[ResizeLargerCorrect](examples/ResizeLargerCorrect.png) |

## Plan

- **Porter/Duff composition**

  - [x] Source
  - [x] Atop
  - [x] Over
  - [x] In
  - [x] Out
  - [x] Dest
  - [x] DestAtop
  - [x] DestOver
  - [x] DestIn
  - [x] DestOut
  - [x] Clear
  - [x] Xor
  - [x] Bonus: Lighter

- **Transformation**

  - [x] Flip horizontally/vertically
  - [x] Rotate
  - [x] Resized

- **Effects**
  - [ ] Blur
  - [ ] Pixelate
