# Go Mux Postgres

## Migration and seeding

```sql
CREATE TABLE IF NOT EXISTS products_tab (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    quantity INT
);

INSERT INTO products_tab (name, description, quantity) VALUES ('Sun - Dried Tomatoes', 'Quisque id justo sit amet sapien dignissim vestibulum.', 160);
INSERT INTO products_tab (name, description, quantity) VALUES ('Beans - Fava, Canned', 'Quisque arcu libero, rutrum ac, lobortis vel, dapibus at, diam. Nam tristique tortor eu pede.', 128);
INSERT INTO products_tab (name, description, quantity) VALUES ('Nantucket Pine Orangebanana', 'Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum.', 133);
INSERT INTO products_tab (name, description, quantity) VALUES ('Water - Aquafina Vitamin', 'Curabitur in libero ut massa volutpat convallis. Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo. Maecenas pulvinar lobortis est.', 140);
INSERT INTO products_tab (name, description, quantity) VALUES ('Caviar - Salmon', 'Suspendisse potenti. Nullam porttitor lacus at turpis.', 72);
INSERT INTO products_tab (name, description, quantity) VALUES ('Nantucket Apple Juice', 'Morbi ut odio. Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo. In blandit ultrices enim.', 63);
INSERT INTO products_tab (name, description, quantity) VALUES ('Coriander - Seed', 'Etiam vel augue. Vestibulum rutrum rutrum neque. Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia.', 37);
INSERT INTO products_tab (name, description, quantity) VALUES ('Chocolate - Mi - Amere Semi', 'Sed accumsan felis. Ut at dolor quis odio consequat varius.', 175);
INSERT INTO products_tab (name, description, quantity) VALUES ('Bread Foccacia Whole', 'Nullam porttitor lacus at turpis. Donec posuere metus vitae ipsum. Aliquam non mauris. Morbi non lectus.', 96);
INSERT INTO products_tab (name, description, quantity) VALUES ('Rosemary - Dry', 'Phasellus in felis.', 92);
INSERT INTO products_tab (name, description, quantity) VALUES ('Calypso - Pineapple Passion', 'Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci. Mauris lacinia sapien quis libero. Nullam sit amet turpis elementum ligula vehicula consequat.', 102);
INSERT INTO products_tab (name, description, quantity) VALUES ('Parasol Pick Stir Stick', 'Morbi a ipsum.', 27);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - White, Schroder And Schyl', 'Sed accumsan felis. Ut at dolor quis odio consequat varius.', 52);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cleaner - Comet', 'Morbi ut odio.', 111);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cheese - Wine', 'Aliquam non mauris. Morbi non lectus. Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus.', 152);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cake - Lemon Chiffon', 'Sed accumsan felis. Ut at dolor quis odio consequat varius. Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi.', 181);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pheasants - Whole', 'In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa.', 33);
INSERT INTO products_tab (name, description, quantity) VALUES ('Soup - Knorr, Chicken Gumbo', 'Aenean lectus. Pellentesque eget nunc. Donec quis orci eget orci vehicula condimentum. Curabitur in libero ut massa volutpat convallis. Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo.', 96);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Spumante Bambino White', 'Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem. Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy.', 63);
INSERT INTO products_tab (name, description, quantity) VALUES ('Ostrich - Fan Fillet', 'Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis.', 171);
INSERT INTO products_tab (name, description, quantity) VALUES ('Steampan Lid', 'Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui. Maecenas tristique, est et tempus semper, est quam pharetra magna, ac consequat metus sapien ut nunc. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris viverra diam vitae quam.', 15);
INSERT INTO products_tab (name, description, quantity) VALUES ('Guava', 'Mauris ullamcorper purus sit amet nulla. Quisque arcu libero, rutrum ac, lobortis vel, dapibus at, diam.', 175);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Taylors Reserve', 'Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien.', 194);
INSERT INTO products_tab (name, description, quantity) VALUES ('Lamb - Leg, Bone In', 'Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum.', 48);
INSERT INTO products_tab (name, description, quantity) VALUES ('Sauce - Black Current, Dry Mix', 'Cras non velit nec nisi vulputate nonummy.', 12);
INSERT INTO products_tab (name, description, quantity) VALUES ('Mushroom - Enoki, Dry', 'Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus.', 140);
INSERT INTO products_tab (name, description, quantity) VALUES ('Mustard Prepared', 'Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus.', 122);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cilantro / Coriander - Fresh', 'Aliquam non mauris.', 89);
INSERT INTO products_tab (name, description, quantity) VALUES ('Goat - Leg', 'Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit.', 26);
INSERT INTO products_tab (name, description, quantity) VALUES ('Tray - Foam, Square 4 - S', 'Nulla tellus. In sagittis dui vel nisl. Duis ac nibh. Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus.', 181);
INSERT INTO products_tab (name, description, quantity) VALUES ('Lamb - Whole Head Off,nz', 'Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est.', 188);
INSERT INTO products_tab (name, description, quantity) VALUES ('Oregano - Fresh', 'Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci. Mauris lacinia sapien quis libero. Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum.', 192);
INSERT INTO products_tab (name, description, quantity) VALUES ('Peach - Fresh', 'Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo.', 176);
INSERT INTO products_tab (name, description, quantity) VALUES ('Beef - Prime Rib Aaa', 'Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla.', 85);
INSERT INTO products_tab (name, description, quantity) VALUES ('Ginger - Pickled', 'Nullam molestie nibh in lectus. Pellentesque at nulla.', 148);
INSERT INTO products_tab (name, description, quantity) VALUES ('Mix - Cocktail Ice Cream', 'Proin eu mi. Nulla ac enim.', 112);
INSERT INTO products_tab (name, description, quantity) VALUES ('Bay Leaf Ground', 'Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero. Nam dui.', 85);
INSERT INTO products_tab (name, description, quantity) VALUES ('Milk - 1%', 'In sagittis dui vel nisl. Duis ac nibh.', 109);
INSERT INTO products_tab (name, description, quantity) VALUES ('Rabbit - Saddles', 'Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus.', 30);
INSERT INTO products_tab (name, description, quantity) VALUES ('Persimmons', 'Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti. Cras in purus eu magna vulputate luctus.', 50);
INSERT INTO products_tab (name, description, quantity) VALUES ('Ecolab Crystal Fusion', 'Morbi a ipsum. Integer a nibh.', 152);
INSERT INTO products_tab (name, description, quantity) VALUES ('Beef Dry Aged Tenderloin Aaa', 'Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis.', 180);
INSERT INTO products_tab (name, description, quantity) VALUES ('Mustard - Dry, Powder', 'Aliquam erat volutpat.', 54);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Merlot Vina Carmen', 'Nulla tempus. Vivamus in felis eu sapien cursus vestibulum. Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem.', 50);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pepper - Scotch Bonnet', 'Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien.', 14);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Fino Tio Pepe Gonzalez', 'Morbi non lectus. Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl.', 89);
INSERT INTO products_tab (name, description, quantity) VALUES ('Longos - Cheese Tortellini', 'Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor. Duis mattis egestas metus.', 186);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pail - 15l White, With Handle', 'Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam.', 65);
INSERT INTO products_tab (name, description, quantity) VALUES ('Energy Drink - Redbull 355ml', 'Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti.', 163);
INSERT INTO products_tab (name, description, quantity) VALUES ('Spinach - Baby', 'Proin at turpis a pede posuere nonummy. Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi.', 198);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Tribal Sauvignon', 'Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus. In sagittis dui vel nisl. Duis ac nibh.', 93);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cake - Cake Sheet Macaroon', 'Mauris lacinia sapien quis libero. Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum.', 117);
INSERT INTO products_tab (name, description, quantity) VALUES ('Carbonated Water - Raspberry', 'Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum. In hac habitasse platea dictumst.', 113);
INSERT INTO products_tab (name, description, quantity) VALUES ('Sambuca - Opal Nera', 'Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede.', 200);
INSERT INTO products_tab (name, description, quantity) VALUES ('Beef Striploin Aaa', 'In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl.', 63);
INSERT INTO products_tab (name, description, quantity) VALUES ('Trout - Rainbow, Fresh', 'Quisque porta volutpat erat. Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero.', 47);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pepper - Red Thai', 'Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus. Phasellus in felis. Donec semper sapien a libero. Nam dui.', 21);
INSERT INTO products_tab (name, description, quantity) VALUES ('Bouq All Italian - Primerba', 'Etiam justo. Etiam pretium iaculis justo.', 111);
INSERT INTO products_tab (name, description, quantity) VALUES ('Mushroom - Oyster, Fresh', 'Morbi a ipsum. Integer a nibh.', 190);
INSERT INTO products_tab (name, description, quantity) VALUES ('Dasheen', 'Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia. Aenean sit amet justo. Morbi ut odio.', 102);
INSERT INTO products_tab (name, description, quantity) VALUES ('Kolrabi', 'Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla. Nunc purus.', 120);
INSERT INTO products_tab (name, description, quantity) VALUES ('Crackers - Graham', 'Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est. Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros.', 53);
INSERT INTO products_tab (name, description, quantity) VALUES ('Salt - Table', 'Nulla facilisi. Cras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque.', 75);
INSERT INTO products_tab (name, description, quantity) VALUES ('Napkin Colour', 'In sagittis dui vel nisl. Duis ac nibh. Fusce lacus purus, aliquet at, feugiat non, pretium quis, lectus.', 41);
INSERT INTO products_tab (name, description, quantity) VALUES ('Flour Pastry Super Fine', 'Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh.', 157);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cake - Miini Cheesecake Cherry', 'Donec ut dolor.', 46);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cheese - Grana Padano', 'Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus.', 45);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Coteaux Du Tricastin Ac', 'Suspendisse accumsan tortor quis turpis.', 168);
INSERT INTO products_tab (name, description, quantity) VALUES ('Juice - Clam, 46 Oz', 'Donec quis orci eget orci vehicula condimentum. Curabitur in libero ut massa volutpat convallis.', 41);
INSERT INTO products_tab (name, description, quantity) VALUES ('Sea Urchin', 'Cras in purus eu magna vulputate luctus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue.', 199);
INSERT INTO products_tab (name, description, quantity) VALUES ('Mousse - Passion Fruit', 'In hac habitasse platea dictumst. Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem. Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat.', 151);
INSERT INTO products_tab (name, description, quantity) VALUES ('Marzipan 50/50', 'Morbi ut odio. Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo. In blandit ultrices enim.', 113);
INSERT INTO products_tab (name, description, quantity) VALUES ('Yukon Jack', 'Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus.', 122);
INSERT INTO products_tab (name, description, quantity) VALUES ('Jam - Raspberry,jar', 'Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl.', 131);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pork - Liver', 'Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vivamus vestibulum sagittis sapien. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.', 143);
INSERT INTO products_tab (name, description, quantity) VALUES ('Water - Tonic', 'Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam vel augue. Vestibulum rutrum rutrum neque.', 27);
INSERT INTO products_tab (name, description, quantity) VALUES ('Bagelers', 'Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem. Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat.', 173);
INSERT INTO products_tab (name, description, quantity) VALUES ('Chicken Breast Halal', 'Vivamus vestibulum sagittis sapien.', 148);
INSERT INTO products_tab (name, description, quantity) VALUES ('Bread - White Mini Epi', 'Proin leo odio, porttitor id, consequat in, consequat ut, nulla.', 136);
INSERT INTO products_tab (name, description, quantity) VALUES ('Cookie Dough - Chocolate Chip', 'Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.', 189);
INSERT INTO products_tab (name, description, quantity) VALUES ('Water - Evian 355 Ml', 'Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui.', 178);
INSERT INTO products_tab (name, description, quantity) VALUES ('Calypso - Pineapple Passion', 'Duis consequat dui nec nisi volutpat eleifend. Donec ut dolor.', 89);
INSERT INTO products_tab (name, description, quantity) VALUES ('Flour - Teff', 'Praesent blandit. Nam nulla.', 195);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pasta - Shells, Medium, Dry', 'Suspendisse potenti. Cras in purus eu magna vulputate luctus.', 194);
INSERT INTO products_tab (name, description, quantity) VALUES ('Lettuce - Frisee', 'Morbi ut odio. Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo. In blandit ultrices enim. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Proin interdum mauris non ligula pellentesque ultrices.', 84);
INSERT INTO products_tab (name, description, quantity) VALUES ('Extract Vanilla Pure', 'Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi. Cras non velit nec nisi vulputate nonummy.', 152);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - White, Schroder And Schyl', 'Nullam porttitor lacus at turpis. Donec posuere metus vitae ipsum. Aliquam non mauris. Morbi non lectus. Aliquam sit amet diam in magna bibendum imperdiet.', 186);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Pinot Grigio Collavini', 'Vestibulum rutrum rutrum neque. Aenean auctor gravida sem.', 126);
INSERT INTO products_tab (name, description, quantity) VALUES ('Greens Mustard', 'In hac habitasse platea dictumst. Maecenas ut massa quis augue luctus tincidunt.', 145);
INSERT INTO products_tab (name, description, quantity) VALUES ('Wine - Sauvignon Blanc', 'Nulla suscipit ligula in lacus. Curabitur at ipsum ac tellus semper interdum. Mauris ullamcorper purus sit amet nulla.', 115);
INSERT INTO products_tab (name, description, quantity) VALUES ('Coffee - Cafe Moreno', 'Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla.', 153);
INSERT INTO products_tab (name, description, quantity) VALUES ('Muffin - Blueberry Individual', 'Morbi non lectus.', 180);
INSERT INTO products_tab (name, description, quantity) VALUES ('Yogurt - Plain', 'Quisque erat eros, viverra eget, congue eget, semper rutrum, nulla.', 102);
INSERT INTO products_tab (name, description, quantity) VALUES ('Tea - Herbal - 6 Asst', 'Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy.', 98);
INSERT INTO products_tab (name, description, quantity) VALUES ('Bread - Rye', 'Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy. Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue.', 163);
INSERT INTO products_tab (name, description, quantity) VALUES ('Pepper - Green Thai', 'Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci.', 167);
INSERT INTO products_tab (name, description, quantity) VALUES ('Soup Campbells', 'Vivamus vestibulum sagittis sapien.', 120);
INSERT INTO products_tab (name, description, quantity) VALUES ('Ocean Spray - Ruby Red', 'Duis consequat dui nec nisi volutpat eleifend. Donec ut dolor. Morbi vel lectus in quam fringilla rhoncus.', 189);
INSERT INTO products_tab (name, description, quantity) VALUES ('Salt - Rock, Course', 'Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus.', 155);
INSERT INTO products_tab (name, description, quantity) VALUES ('Sugar - White Packet', 'Vestibulum rutrum rutrum neque.', 91);
```

## Exercises

1. Refactor `DELETE /products/:id` to return the id of the deleted product. Send `http.StatusNotFound` response to the client when the product is not found.
2. Refactor `UPDATE /products/:id` to return the newly updated data. Send `http.StatusNotFound` response to the client when the product is not found.
3. Refactor your app to use environment variables instead.
4. Add search by name feature to `GET /products`. Add search to the request query.
5. Add pagination feature to `GET /products`. Add page and size to the request query. When user does not set page or size value, they will be set to their default value. Default value for page is **1** and default value for size is **10**.
6. Add sort by name feature to `GET /products`. Add sort to the request query. If user does not set sort value, it will be set to **asc** by default.
7. Combine point 4 to 6.
