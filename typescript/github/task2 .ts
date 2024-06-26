class Pizza {
    hasPepperoni: boolean;
    hasSauce: boolean;
    hasCheese: boolean;
    bakeTime = 30;
  
    constructor(hasPepperoni: boolean, hasSauce: boolean, hasCheese: boolean) {
      this.hasPepperoni = hasPepperoni;
      this.hasSauce = hasSauce;
      this.hasCheese = hasCheese;
    }
  
    bake():void {
      setTimeout(() => console.log('Enjoy your pizza!'));
    }
  }
  
  class Pasta {
    pastaType: string;
    hasSauce: boolean;
    hasCheese: boolean;
    cookTime = 30;
  
    constructor(pastaType: string, hasSauce: boolean, hasCheese: boolean) {
      this.pastaType = pastaType;
      this.hasCheese = hasCheese;
      this.hasSauce = hasSauce;
    }
  
    cook():void {
      setTimeout(() => console.log('Enjoy your pasta!'));
    }
  }
  
  class Kitchen {
    makeDish(dish: Pizza | Pasta):string {
      if ('bake' in dish) {
        dish.bake();
        return `Enjoy your ${dish.constructor.name} with ${dish.hasPepperoni}, ${dish.hasCheese}, ${dish.hasSauce}`;
      } else if ('cook' in dish) {
        dish.cook();
        return `Enjoy your ${dish.constructor.name}`;
      }
      throw new Error('Unknown dish');
    }
  }
  
  const kitchen = new Kitchen();
  const pizza = new Pizza(true, true, false);
  const pasta = new Pasta('spagetti', true, true);
  
  kitchen.makeDish(pizza);
  kitchen.makeDish(pasta);