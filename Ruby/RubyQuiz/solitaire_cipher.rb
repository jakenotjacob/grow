#!/usr/bin/env ruby
#
#Solitaire Cypher
# - encrypyption and decryption of

str = "hello, there world!"


###ENCRYPTION###
#strip any non A characters and upcase
str.scan(/\w/).map(&:upcase)

#Use Solitair to generate a keystream letter for each letter in message

#Convert message from step 1 into numbers (a=1, b=2,...)

#Convert keystream letters via same methos (a=1, b=2,...)

#Add numbers from above steps, subtract 26 from result if greater than 26

#Convert back to letters



###DECRYPTION###




class Solitaire
  attr_accessor :message, :keyed
  def initialize(message)
    @message = sanitize(message)
    @deck = [*(1..52), 'A', 'B']
  end

  def encrypt
    #(1)sanitize()
    #(2)generate keystream
    #(3)convert original message to numbers
    #(4)convert keystream letters to numbers
    #(5)Combine (3) and (4) to new array.
    ###Subtract 26 from result if greater than 26
    #(6)Convert (5) back to letters
  end

  #TODO - first element cannot be changed
  def move_down(card)
    deck = @deck.dup
    card_index = @deck.index(card)
    if card_index == (deck.length-1)
      deck[card_index] = @deck[1]
      deck[1] = card
    else
      deck[card_index] = @deck[card_index+1]
      deck[card_index+1] = card
    end
    @deck = deck
  end

  def move_up(card)
    deck = @deck.dup
    card_index = @deck.index(card)
    if card_index == 1
      deck[0] = @deck[-1]
      deck[-1] = card
    else
      deck[card_index] = @deck[card_index-1]
      deck[card_index-1] = card
    end
    @deck = deck
  end

  def gen_keystream
    deck = @deck.dup
    deck.move_down('B')
    deck.move_down('B')
  end

  def sanitize(message)
    fill_count = (5-(message.length % 5))%5
    msg = message.scan(/\w/).map(&:upcase)
    msg << "X"*fill_count
    msg.reject(&:empty?)
    msg.join.scan(/.{5}/)
  end

end



