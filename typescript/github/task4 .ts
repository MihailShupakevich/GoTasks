
interface IBaseProductParams {
  name: string;
  price: number;
  discount: number;
}
class Shop {
  private items: BaseProduct[];

  constructor(items: BaseProduct[] = []) {
    this.items = items;
  }

  public addItem(item: BaseProduct) {
    this.items.push(item);
  }

  public getItems(): BaseProduct[] {
    return this.items;
  }
}
class BaseProduct {
  name: string;
  price: number;
  discount: number;

  constructor(params: IBaseProductParams) {
    this.name = params.name;
    this.price = params.price;
    this.discount = params.discount;
  }
}

class Banan extends BaseProduct {
  size: string;
  constructor(price: number, discount: number, size: string) {
    super({ name: 'Banan', price, discount });
    this.size = size;
  }
}
class IceCream extends BaseProduct {
  withGlace: boolean;

  constructor(price: number, discount: number, withGlace: boolean) {
    super({ name: 'IceCream', price, discount });
    this.withGlace = withGlace;
  }
}
const shop = new Shop([new BaseProduct({name:'123',discount: 11,price: 11})]);
const iceCream = new IceCream(10, 0, true);
const banan = new Banan(5, 0.1, 'small');
shop.addItem(banan)
