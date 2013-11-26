import unittest

from engine import player, consts, types


class PlayerTest(unittest.TestCase):
    def testPlayerCreate(self):
        player_ = player.Player(1)
        self.assertEqual(player_.player_id, 1)
        self.assertEqual(player_.units, {})

    def testPlayerAddUnit(self):
        player_ = player.Player(1)
        tank = types.new_unit('Tank', consts.RED)
        player_.add_unit(tank, 1)
        self.assertEqual(player_.units, {1: tank})

    def testPlayerGetUnit(self):
        player_ = player.Player(1)
        tank = types.new_unit('Tank', consts.RED)
        player_.add_unit(tank, 1)
        returned_tank = player_.get_unit(1)
        self.assertEqual(returned_tank, tank)
        self.assertRaises(Exception, player_.get_unit, 2)

    def testPlayerSerialize(self):
        player_ = player.Player(1)
        self.assertEqual(player_.serialize(True), {})

if __name__ == '__main__':
    unittest.main()
